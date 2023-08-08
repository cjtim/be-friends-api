# Builder
FROM golang:1.17-alpine as builder
WORKDIR $GOPATH/src/github.com/cjtim/be-friends-api

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod tidy
RUN go mod download

COPY . .

ARG GOARCH
ARG GOOS=linux
ARG CGO_ENABLED=0

RUN go build -o main main.go

# Production container
FROM alpine:3.18.3  
RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    libc6-compat
WORKDIR /root/
COPY --from=builder /go/src/github.com/cjtim/be-friends-api .

EXPOSE 8080

ARG TAG
ENV VERSION=${TAG}

CMD ["/root/main"]