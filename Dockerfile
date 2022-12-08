FROM golang:alpine AS builder

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

FROM alpine
#更新Alpine的软件源为国内源，提高下载速度
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories

RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash
# 设置时区为上海
RUN apk -U add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

COPY wait-for-it.sh .
COPY configs.yaml .

COPY --from=builder /build/bbs .


RUN chmod 755 wait-for-it.sh



