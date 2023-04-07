#!/bin/sh

docker stop $MONGODB_PID
kill -9 $API_PID
fuser -k 5000/tcp
kill -9 $WEB_PID

cd ..