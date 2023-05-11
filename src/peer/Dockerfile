FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /app

COPY ./src/peer/go.mod ./

RUN go mod download

CMD ["air", "-c", ".air.toml"]