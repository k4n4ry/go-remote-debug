FROM golang:1.16.5

WORKDIR /app

COPY . .

RUN go get -u github.com/go-delve/delve/cmd/dlv 
