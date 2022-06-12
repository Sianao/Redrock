FROM golang:1.17 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

ENV GOPROXY=https://goproxy.cn,direct
WORKDIR "/Dir"
MAINTAINER  Sianao
COPY . .

RUN go mod download
COPY . .
RUN go build -o app .
FROM busybox
COPY --from=bui /Dir/app /

ENTRYPOINT ["/app"]
