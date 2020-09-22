FROM golang:1.15-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache alpine-sdk mysql-client mysql-dev build-base tzdata bash zsh vim less

RUN go get github.com/rubenv/sql-migrate/...