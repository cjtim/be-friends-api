FROM golang:1.17.6-alpine as builder
WORKDIR $GOPATH/src/github.com/cjtim/be-friends-api

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod tidy

COPY . .

ARG GOARCH
ARG GOOS=linux
ARG CGO_ENABLED=0

RUN go run github.com/prisma/prisma-client-go generate
# Build the binary.
RUN go build -o main main.go

FROM alpine:latest  
RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    libc6-compat
WORKDIR /root/
COPY --from=builder /go/src/github.com/cjtim/be-friends-api .
EXPOSE 8080
CMD ["/root/main"]