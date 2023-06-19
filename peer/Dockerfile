FROM docker.uclv.cu/golang:1.20-alpine as builder

WORKDIR /build

COPY ./ ./

RUN go build -o main

FROM docker.uclv.cu/redis:latest

WORKDIR /peer

COPY --from=builder /build/main ./

COPY ./cert ./cert

COPY ./start.sh ./

# EXPOSE 8080

CMD [ "bash", "start.sh" ]

# FROM docker.uclv.cu/golang:1.20-alpine

# WORKDIR /peer

# COPY ./ ./

# # EXPOSE 8080
# # EXPOSE 32140/udp

# CMD [ "go", "run", "main.go" ]

# RUN go install github.com/cosmtrek/air@v1.43.0

# RUN go clean -modcache

# COPY ./src/peer/go.mod ./

# RUN go mod download

# CMD ["air", "-c", ".air.toml"]