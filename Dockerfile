FROM golang:1.17.3

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY ./main.go  ./

EXPOSE $PORT

CMD go run main.go
