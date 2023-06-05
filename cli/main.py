import random
import docker
import sys
import os
from pprint import pprint
from ipaddress import IPv4Network


client = docker.from_env()


def getContainers():
    containers = client.networks.get("gotify-net").attrs['Containers']
    containers = [(id[:12], containers[id]['IPv4Address'].split('/')[0]) for id in containers]
    return containers


def getMongoDbNets(id: str):
    ips = [container for container in client.containers.
        list(filters={
            "id": id,
            "ancestor": "docker.uclv.cu/mongo:latest"
        })]

    ipAddress = ips[0].attrs['NetworkSettings']['Networks']['gotify-net']['IPAddress']

    hostIp = ips[0].attrs['HostConfig']['PortBindings']['27017/tcp'][0]['HostIp']
    hostPort = ips[0].attrs['HostConfig']['PortBindings']['27017/tcp'][0]['HostPort']

    return (ipAddress, (hostIp, int(hostPort)))


def getAvailableIP(subnet: str, excluded: list = []):
    subnet = IPv4Network(subnet)
    for ip in subnet.hosts():
        if str(ip) not in excluded and not str(ip).endswith('0.1'):
            return str(ip)
    return "0.0.0.0"


def getDockerIPAvailable():
    ips = [ip for _, ip in getContainers()]
    return getAvailableIP("192.168.0.0/16", ips)


def getBindings():
    bindings = []

    for id, _ in getContainers():
        container = client.containers.get(id)
        ports = container.attrs['NetworkSettings']['Ports']

        for port in ports:
            if ports[port] != None:
                for binding in ports[port]:
                    bindings.append((id, binding['HostIp'], int(binding['HostPort'])))

    return bindings


def getHostIPAvailable():
    ips = [ip for _, ip, _ in getBindings()]
    return getAvailableIP("127.0.0.0/16", ips)


def run_container(image: str, ip: str = '0.0.0.0', env: dict = None, vol: list = None, ports: dict= None):
    def d(d: dict):
        if d == None:
            return {}
        return d

    def l(l: list):
        if l == None:
            return []
        return l

    vol = l(vol)
    env = d(env)
    ports = d(ports)
    
    container = client.containers.run(
        image=image, 
        detach=True,
        auto_remove=True,
        init=True,
        network="gotify-net",
        hostname=ip,
        volumes=vol,
        ports=ports,
        environment=env
    )
    return container


def rm_rand_containers(image:str, amount: int):
    """Remove `amount` docker containers with `image` as the base image"""

    containers = client.containers.list(filters={
        "ancestor": f"gotify-{image}:latest"
    })
    containers = random.sample(containers, amount)
    
    for container in containers:
        try:
            container.stop()
            container.kill()
        except:...


if __name__ == '__main__':
    
    command = sys.argv[1]

    try:
        image = sys.argv[2]
    except:
        image = ""

    if command == "start":

        dns = run_container(
            image='dns'
        )
        
        dockerIp = getDockerIPAvailable()
        db = run_container(
            image='mongo',
            ip=dockerIp,
            env={
                'MONGO_INITDB_ROOT_USERNAME': 'user',
                'MONGO_INITDB_ROOT_PASSWORD': 'password'
            },
        )
        peer = run_container(
            image='peer',
            env={
                'MONGODB_IP': dockerIp
            },
            # vol=['/home/leandro/go/src/github.com/science-engineering-art/gotify/data:/data/db']
        )

        dockerIp = getDockerIPAvailable()
        db = run_container(
            image='mongo',
            ip=dockerIp,
            env={
                'MONGO_INITDB_ROOT_USERNAME': 'user',
                'MONGO_INITDB_ROOT_PASSWORD': 'password'
            },
        )
        api = run_container(
            image='api',
            env={
                'MONGODB_IP': dockerIp
            },
        )

        web = run_container(
            image='web'
        )

    elif command == "kill":

        if image == "all" and len(sys.argv) == 3:
            for container in client.containers.list(all=True):
                try:
                    container.stop()
                    container.remove()
                except:...

        elif len(sys.argv) == 3:
            # remove a specific container 
            id = sys.argv[2]
            container = client.containers.get(id)
            
            try:
                container.stop()
                container.remove()
            except:...

        elif image in ["dns", "web", "api", "peer"] \
            and len(sys.argv) == 4:
            
            # amount of containers to trash
            amount = int(sys.argv[3])
            
            # remove N-1 containers
            if amount == -1:
                containers = client.containers.list(filters={
                    "ancestor": f"{image}:latest"
                })
                rm_rand_containers(image, len(containers) - 1)
            
            # delete `amount` 
            elif amount > 0:
                rm_rand_containers(image, amount)

    elif command == "run":
        try:
            # amount of container to run
            amount = int(sys.argv[3])
            
            for i in range(amount):
                client.containers.run(
                    image=image, 
                    detach=True,
                    auto_remove=True,
                    init=True,
                    network="gotify-net",
                    hostname='0.0.0.0',
                )
        except:
            # run a container with specific IP
            ip = sys.argv[3]

            client.containers.run(
                image=image, 
                detach=True,
                auto_remove=True,
                init=True,
                network="gotify-net",
                hostname=ip,
            )

    elif command == "list":

        if image in ["dns", "web", "api", "peer"]:
            containers = client.containers.list(filters={
                "ancestor": f"{image}:latest"
            })
            for container in containers:
                print(container)

    elif command == "build":

        for container in client.containers.list():
            container.stop()

        for img in ["dns", "web", "api", "peer"]:
            client.images.build(
                path=f"../{img}",
                dockerfile="Dockerfile",
                rm=False,
                tag=f"{img}:latest"
            )
        
        old_imgs = [img for img in client.images.list() if len(img.attrs['RepoTags']) == 0]

        for img in old_imgs:
            client.images.remove(image=img.id)

        for container in client.containers.list(all=True):
            try:
                container.stop()
                container.remove()
            except:...
