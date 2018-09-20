FROM golang:alpine as builder
RUN apk add --no-cache git build-base && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/main" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories && \
    echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories && \
    apk add --no-cache upx

ENV GO111MODULE=on
WORKDIR /go/src/github.com/AlphaWong/log-entry
COPY . .

RUN  CGO_ENABLE=0 GOOS=linux go build \
 -tags netgo \
 -installsuffix netgo,cgo \
 -v -a \
 -ldflags '-s -w -extldflags "-static"' \
 -o app

RUN upx -7 -qq app && \
  upx -t app && \
  mv ./app /go/bin/app

FROM gcr.io/distroless/base
COPY --from=builder /go/bin/app /
COPY --from=builder /go/src/github.com/AlphaWong/log-entry/config/config.yaml /config/config.yaml
ENTRYPOINT ["./app"]
EXPOSE 80