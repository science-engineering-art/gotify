FROM alpine:latest

WORKDIR /models

COPY ./ ./

ENTRYPOINT [ "/bin/sh" ]