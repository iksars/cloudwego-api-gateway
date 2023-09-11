#!/bin/bash

frontEndPath=/home/iksar/goproject/src/github.com/iksars/cloudwego-api-gateway/pkg/IDL-Management/vue-front
backEndPath=/home/iksar/goproject/src/github.com/iksars/cloudwego-api-gateway/pkg/IDL-Management/hertz-back
originPath=/home/iksar/goproject/src/github.com/iksars/cloudwego-api-gateway

gnome-terminal -t "backend" -- bash -c "cd ${backEndPath}; go build -o hertz-back; ./hertz-back; exec bash"
sleep 1
gnome-terminal -t "frontend" -- bash -c "cd ${frontEndPath}; npm run dev; exec bash"
sleep 1
gnome-terminal -t "agw" -- bash -c "cd ${originPath}; go build; ./cloudwego-api-gateway ; exec bash"



