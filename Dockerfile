FROM golang:1.14

WORKDIR /go/src/

COPY . .

RUN go run main.go


