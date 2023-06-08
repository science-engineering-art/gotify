import requests
import random
import docker
import sys
import time

images = ["web", "api", "dns", "peer", "trackers"]
client = docker.from_env()
url = 'https://gotify.com'

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

def create_k_containers(image, k):
    l = []
    for _ in range(0, k):
        l.append(run_container(image,"0.0.0.0"))

    return l

red = []

def main():
    # Create K containers for images
    list_dns = create_k_containers(images[0], 10)
    list_web = create_k_containers(images[1], 10)
    list_api = create_k_containers(images[2], 10)
    list_peer = create_k_containers(images[3], 10)
    list_trackers = create_k_containers(images[4], 10)

    red.append(list_web)
    red.append(list_api)
    red.append(list_dns)
    red.append(list_peer)
    red.append(list_trackers)

    # Request of clients
    t = 0
    id_node = 0
    while True:
        id_node = rand_url()

        payload = {'key1': 'value1', 'key2': 'value2'} # estructurar datos .json de la musica
        response = requests.get(id_node, params=payload)
        print(response.url)

        time.sleep(10)
        t += 10

        payload = {'key1': 'value1', 'key2': 'value2'} # estructurar datos .json de la musica
        response = requests.post(id_node, params=payload)
        print(response.url)

        time.sleep(10)
        t += 10

        payload = {'key1': 'value1', 'key2': 'value2'} # estructurar datos .json de la musica
        response = requests.delete(id_node, params=payload)
        print(response.url)
        t += 10

        # Unir y desconectar contenedores de la red

        if t % 50 == 0:
            r = random.Random().randint(1,5)
            Disconecter(r)

        if t % 60 == 0:
            r = random.Random().randint(1,5)
            Connecter(r)


def Disconecter(k):
    for i in range(0,k):
        r = random.Random().randint(0,5)
        s = random.Random().randint(0,len(red[r]))

        container = red[r][s]
        #id = container
        #container = client.containers.get(id)
            
        try:
            container.stop()
            container.remove()
        except:
            Exception()
        

    

def Connecter (k):
    for i in range(0,k):
        r = random.Random().randint(0,5)
        
        container = run_container(images,"0.0.0.0")
        red[r].append(container)
        #id = container
        #container = client.containers.get(id)

        ip_address = container.attrs['NetworkSettings']['IPAddress']
        print('Nuevo contenedor ', images[r], ' conectado en el ip: ',ip_address)

def rand_url():
    r = random.Random().randint(0,len(red[1])) # Escoger una api random

    container = red[1][r]
    ip_address = container.attrs['NetworkSettings']['IPAddress']
    
    return 'http\\:' + ip_address




