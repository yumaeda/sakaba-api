ARG GO_VERSION=1.17.3

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY ./main.go  ./

CMD ["sh", "-c", "go run main.go"]
