FROM golang:1.22

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY *.go ./

RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 3333

CMD ["app"]