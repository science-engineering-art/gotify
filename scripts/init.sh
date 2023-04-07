#!/bin/sh

source mongo.sh

cd ../src/api
go run main.go &

export API_PID=$!

cd ../client
npm run dev &

export WEB_PID=$!

cd ../..