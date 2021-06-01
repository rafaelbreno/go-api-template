FROM golang:1.16-alpine

WORKDIR /app

RUN apk update \
  && apk add --virtual build-deps gcc git curl tar \
  && rm -rf /var/cache/apk/*

RUN addgroup -S golang \
  && adduser -S -G golang golang

COPY . .

RUN go build main.go && ./main

EXPOSE 8080
CMD ["mailroom"]
