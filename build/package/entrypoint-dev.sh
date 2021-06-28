#!/bin/sh

#todo[petr]: remove dependencies downloading step
#echo Downloading dependencies
#go get -t ./...

echo Building application
cd cmd/console && GOOS=linux GOARCH=amd64 go build -race -o ~/application || exit

#echo Copying config
#cp /go/src/github.com/vehsamrak/project/configs/config-dev.yml ~/config.yml

echo Running application
cd ~ && ./application
