FROM spotify-rpc:latest

WORKDIR /app/peer

COPY ./src/peer/go.mod ./

RUN go mod download

COPY ./src/docker/peer/go.work ../

CMD ["air", "-c", ".air.toml"]