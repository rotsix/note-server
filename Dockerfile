# building
FROM golang:alpine as builder

RUN apk update && apk add --update git gcc
RUN mkdir -p /go/src/note-server
ENV GOPATH /go
WORKDIR /go/src/note-server
COPY go.mod go.sum ./
RUN go mod download
# RUN go test -v ./...
COPY . .
RUN go build -v -o ./server ./cmd/server/main.go


# running
FROM alpine

LABEL maintainer="Victor Franzi <victor.franzi@gmail.com>"

RUN adduser -S -D -H -h /app user
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /go/src/note-server/server .
COPY --chown=user:users ./configs .
ENV CONF_LOCATION /app/config.json
ENV CERT_LOCATION /app/cert.pem
ENV KEY_LOCATION  /app/key.pem
EXPOSE 8080
USER user

CMD ["./server"]
