# Simple Golang HTTP Server

## Docker instructions

### How to build docker image

`docker build -t my-golang-httpserver .`

### How to run docker container

```
docker run -p 3333:3333 -t --rm --name my-running-app my-golang-httpserver
```