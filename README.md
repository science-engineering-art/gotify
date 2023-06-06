# Gotify

Spotify implementation in Golang.

## Execution of services

From the root directory execute the following commands.

### Turn on

DNS: 192.168.0.2

```sh
python main.py build
python main.py start
```

```sh
docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 gotify-net
```

```sh
python cli/main.py up
```

### Turn off

```sh
python cli/main.py down
```
