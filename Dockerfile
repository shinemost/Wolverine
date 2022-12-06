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

RUN go mod download

COPY . .

RUN go build -o bbs .

FROM alpine:3.17.0


COPY configs.yaml .
COPY --from=builder /build/bbs /

ENTRYPOINT ["/bbs","configs.yaml"]

EXPOSE 8071

CMD ["/bbs"]

