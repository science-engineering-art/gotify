<div align="center"> 
    <img height="300" src="./docs/gopher.svg"> 
    <h1> Gotify </h1>
    <p> Spotify implementation in Golang </p>
</div>


## Documentation

Go to file: [gotify.md](https://github.com/science-engineering-art/gotify/blob/master/docs/gotify.md)

## Kademlia Package

Take a look at our implementation of the [Kademlia Protocol](https://github.com/science-engineering-art/kademlia-grpc)

## Execution of services

### Set a Docker network interface

```sh
docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 gotify-net
```

Then configure your DNS by adding nameserver: 192.168.0.2

### Turn on

```sh
python cli/main.py build
python cli/main.py up
```

or

```sh
python cli/main.py rebuild
```

### Turn off

```sh
python cli/main.py down
```
