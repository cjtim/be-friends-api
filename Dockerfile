FROM golang:1.17.6-buster as builder
WORKDIR $GOPATH/src/github.com/cjtim/be-friends-api

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod tidy
RUN go mod download

COPY . .

ARG GOARCH
ARG GOOS=linux
ARG CGO_ENABLED=0

RUN go run github.com/prisma/prisma-client-go generate

RUN go build -o main main.go

FROM alpine:latest  
RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    libc6-compat \
    openssl
WORKDIR /root/
COPY --from=builder /go/src/github.com/cjtim/be-friends-api .
EXPOSE 8080
CMD ["/root/main"]