import sys
import os
from pprint import pprint
from container import *


"""
python main.py [command] [args]

                rebuild -- elimina las imagenes anteriores, prepara las imagenes e cada tipo
                de node, luego levanta la red con un nodo de cada tipo
                
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

        for img in ["dns", "web", "api", "peer", "tracker"]:
            client.images.build(
                path=f"../{img}",
                dockerfile="Dockerfile",
                rm=False,
                tag=f"{img}:latest"
            )
        
        old_imgs = [img for img in client.images.list() if len(img.attrs['RepoTags']) == 0]

        for img in old_imgs:
            try:
                client.images.remove(image=img.id)
            except: ...

        for container in client.containers.list(all=True):
            try:
                container.stop()
                container.remove()
            except:...
        
        os.system("docker rmi $(docker images | grep '<none>' | awk '{print $3}')")

    if command == "rebuild":
        
        print("\nRebuild all dependencies\n")
        os.system("cd .. && make vendor && cd cli")
        
        print("\nRebuild Docker Images\n")
        os.system("python3 main.py build")
        
        print("\nNetworking UP\n")
        os.system("python3 main.py up")
        os.system("docker ps -a")

    elif command == "up":

        dns = run_container(image='dns')
        peer = run_container(image='peer')
        tracker = run_container(image='tracker')
        web = run_container(image='web')
        api = run_container(image='api')

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

        elif image in ["dns", "web", "api", "peer", "tracker"] \
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

        if image in ["dns", "web", "api", "peer", "tracker"]:
            containers = client.containers.list(filters={
                "ancestor": f"{image}:latest"
            })
            for container in containers:
                print(container)
