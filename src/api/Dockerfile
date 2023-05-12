FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /api

RUN go install github.com/cosmtrek/air@v1.43.0

RUN go clean -modcache

COPY ./src/api/go.mod ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]