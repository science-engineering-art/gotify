import sys
from pprint import pprint
from container import *


"""
python main.py [command] [args]
                
                build -- construye las imágenes de Docker correspondientes
                    a cada node de la red

                up -- levanta la red con un nodo de cada tipo 
                    DNS, WEB, API y PEER (con sus respectivas DB)  
                
                down -- cierra todos los nodos activos de la red
                
                run <node> <IP> -- levanta un node con un IP específico

                run <node> <amount> -- corre `amount` nodos del tipo
                    especificado
                
                kill <id> -- elimina el node con el mismo ID que 
                    el que fue pasado
                    
                kill <node> <amount> -- elimina `amount` nodos random 
                    del tipo de nodo especificado (`node`)

                list <node> -- muestra la lista de nodos activos del 
                    tipo especificado
"""


if __name__ == '__main__':

    command = sys.argv[1]

    try:
        image = sys.argv[2]
    except:
        image = ""

    if command == "build":

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

    elif command == "up":

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

    elif command == "down":
        for container in client.containers.list(all=True):
                try:
                    container.stop()
                    container.remove()
                except:...

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

    elif command == "kill":

        if len(sys.argv) == 3:
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

    elif command == "list":

        if image in ["dns", "web", "api", "peer"]:
            containers = client.containers.list(filters={
                "ancestor": f"{image}:latest"
            })
            for container in containers:
                print(container)
