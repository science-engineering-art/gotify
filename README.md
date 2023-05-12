# Spotify

Spotify implementation in Golang.

## Execution of services

From the root directory execute the following commands.

### Turn on

```sh
docker run --rm -d --name mongodb -p 127.0.0.1:27017:27017 -e MONGO_INITDB_ROOT_USERNAME=user -e MONGO_INITDB_ROOT_PASSWORD=password docker.uclv.cu/mongo
```

From the `src/api` and `src/peer` execute `air` respectivaly.

### Turn off

```sh
docker stop $(docker ps -a -q)
```
