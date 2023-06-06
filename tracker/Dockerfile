FROM docker.uclv.cu/golang:1.20-alpine

WORKDIR /peer

COPY ./ ./

EXPOSE 8081
EXPOSE 32141/udp

CMD [ "go", "run", "main.go" ]