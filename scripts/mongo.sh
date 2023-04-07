#!/bin/sh

MONGODB_PID=$(docker run --rm --name mongodb -d -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=user -e MONGO_INITDB_ROOT_PASSWORD=pass docker.uclv.cu/mongo)