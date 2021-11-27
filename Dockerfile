FROM golang:1.17.3

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY src ./src

RUN go build -tags=nomsgpack -o ./app ./src/main.go

EXPOSE $PORT

ENTRYPOINT ["./app"]
