FROM golang:1.17.3

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY infrastructure ./infrastructure

RUN go build -o ./app ./main.go

EXPOSE $PORT

ENTRYPOINT ["./app"]
