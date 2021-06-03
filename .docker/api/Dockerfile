FROM golang:1.16-alpine

WORKDIR /app

RUN apk update \
  && apk add --virtual build-deps gcc git curl tar \
  && rm -rf /var/cache/apk/*

RUN addgroup -S golang \
  && adduser -S -G golang golang

COPY services/api .

RUN go build main.go
EXPOSE 8070

CMD ["./main"]
