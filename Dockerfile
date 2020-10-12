FROM golang:1.15-alpine AS before

RUN apk update && \
    apk upgrade && \
    apk add --no-cache alpine-sdk mysql-client mysql-dev build-base tzdata bash zsh vim less

RUN go get github.com/rubenv/sql-migrate/...

WORKDIR /app
copy . .
RUN go build server.go

FROM alpine:latest AS after
WORKDIR /root/
COPY --from=0 /app/server .
CMD ["./server"]