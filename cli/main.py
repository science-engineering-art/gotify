import random
import docker
import sys


client = docker.from_env()


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
    image = sys.argv[2]
    
    if command == "kill":

        if image == "all" and len(sys.argv) == 3:
            for container in client.containers.list():
                try:
                    container.stop()
                    container.kill()
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
                    "ancestor": f"gotify-{image}:latest"
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
                    image=f'gotify-{image}', 
                    detach=True,
                    auto_remove=True,
                    init=True,
                    network="gotify_default",
                    hostname='0.0.0.0',
                )
        except:
            # run a container with specific IP
            ip = sys.argv[3]

            client.containers.run(
                image=f'gotify-{image}', 
                detach=True,
                auto_remove=True,
                init=True,
                network="gotify_default",
                hostname=ip,
            )

    elif command == "list":

        if image in ["dns", "web", "api", "peer"]:
            containers = client.containers.list(filters={
                "ancestor": f"gotify-{image}:latest"
            })
            for container in containers:
                print(container)
