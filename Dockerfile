FROM golang:alpine AS builder

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY src ./src

RUN go build -ldflags "-s -w" -tags=nomsgpack -o server ./src/main.go

EXPOSE $PORT

FROM golang:alpine as runner
COPY --from=builder /go/src/server /opt/app/
ENTRYPOINT ["/opt/app/server"]
