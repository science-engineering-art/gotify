FROM spotify-rpc:latest

WORKDIR /app/api

COPY ./src/api/go.mod ./

RUN go mod download

COPY ./src/docker/api/go.work ../

CMD ["air", "-c", ".air.toml"]