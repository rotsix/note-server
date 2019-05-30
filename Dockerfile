# building
FROM golang:alpine as builder

RUN apk update && apk add git
RUN mkdir -p /go/src/server
COPY . /go/src/server
WORKDIR /go/src/server
RUN go get -d ./...
RUN go build -o ./server cmd/server/main.go


# running
FROM alpine

LABEL maintainer="Victor Franzi <victor.franzi@gmail.com>"

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /go/src/server/server .
COPY ./configs .
ENV CONF_LOCATION /app/config.json
EXPOSE 8080

CMD ["./server"]
