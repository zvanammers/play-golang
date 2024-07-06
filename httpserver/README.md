# Simple Golang HTTP Server

https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go

## Docker instructions

### How to build docker image

`docker build -t my-golang-httpserver .`

### How to run docker container

```
docker run -p 3333:3333 -t --rm --name my-running-app my-golang-httpserver
```

## CURL Commands

### POST Form Data

`curl -v -X POST -F 'name=Zoe' 'http://localhost:3333/hello'`