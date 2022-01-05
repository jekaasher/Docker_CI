# syntax=docker/dockerfile:1
FROM golang:rc-buster AS builder
WORKDIR /app
COPY app/go.mod    ./
COPY app/go.sum    ./
RUN go mod download
COPY app/*.go ./
RUN go build -o main
RUN apt-get -y update

FROM golang:rc-alpine  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main ./
CMD ["./main"]  
