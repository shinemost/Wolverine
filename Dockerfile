FROM golang:1.16-alpine3.13 AS builder

LABEL maintainer="hjfu"

ENV GO11MODULE=on \
    CGO_ENABLE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go build -mod=mod -o bbs .

FROM alpine:3.13

COPY sql/bbs/init.sql .
COPY wait-for.sh .
COPY configs.yaml .

COPY --from=builder /build/bbs .

RUN chmod 755 wait-for.sh bbs



