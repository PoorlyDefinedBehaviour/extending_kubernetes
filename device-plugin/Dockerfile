### builder

FROM golang:1.19-alpine as builder

ADD . /app

WORKDIR /app/cmd

RUN go build -v

### runner

FROM alpine:latest

COPY --from=builder /app/cmd/cmd /usr/local/bin/k8s-device-plugin

CMD ["k8s-device-plugin"]
