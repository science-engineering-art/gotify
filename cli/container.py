import random
import docker
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
    
    if ip == '0.0.0.0':
        ip = getDockerIPAvailable()
    
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
