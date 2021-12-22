FROM golang:alpine AS builder

RUN apk add --no-cache upx

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download
COPY src ./src

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -tags=nomsgpack -o server ./src/main.go && \
    upx --best server -o _upx_server && \
    mv -f _upx_server server

EXPOSE $PORT

FROM scratch as runner
COPY --from=builder /go/src/server /opt/app/
ENTRYPOINT ["/opt/app/server"]
