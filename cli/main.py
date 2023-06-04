import random
import docker
import sys


images = ["dns", "web", "api", "peer"]
client = docker.from_env()


def run_container(image: str, ip: str):
    container = client.containers.run(
        image=f'gotify-{image}', 
        detach=True,
        auto_remove=True,
        init=True,
        network="gotify_default",
        hostname=ip,
    )
    return container


def rm_rand_containers(image:str, amount: int):
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
    if sys.argv[1] == "kill":
        if sys.argv[2] == "all" and len(sys.argv) == 3:
            for container in client.containers.list():
                try:
                    container.stop()
                    container.kill()
                except:...

        elif len(sys.argv) == 3:
            container = client.containers.get(sys.argv[2])
            try:
                container.stop()
                container.remove()
            except:...

        elif sys.argv[2] in images:
            if len(sys.argv) == 4 and sys.argv[3] != "-1":
                rm_rand_containers(sys.argv[2], int(sys.argv[3]))
            elif type(sys.argv[3]) == str:
                ...

    elif sys.argv[1] == "run":
        try:
            for i in range(int(sys.argv[3])):
                run_container(sys.argv[2], '0.0.0.0')
        except:
            run_container(sys.argv[2], sys.argv[3])
