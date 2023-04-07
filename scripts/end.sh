#!/bin/bash

docker stop $MONGODB_PID

kill -9 $API_PID
fuser -k 5000/tcp

list=($(pstree -A -p $WEB_PID | grep -Eow "[0-9]+"))

kill -9 ${list[0]}
kill -9 ${list[1]}
