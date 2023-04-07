#!/bin/sh

source mongo.sh

cd ../src/api
nohup go run main.go 2>/dev/null &

export API_PID=$!

cd ../client
nohup npm run dev 2>/dev/null &

export WEB_PID=$!

cd ../../scripts
