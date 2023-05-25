FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /kademlia

RUN go install github.com/cosmtrek/air@v1.43.0

RUN go clean -modcache

COPY ./src/kademlia/go.mod ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]