### Builder

FROM golang:1.19-alpine as builder

COPY . /app

WORKDIR /app/cmd

RUN go build -v -ldflags "-s -w"

### Runner

FROM alpine:latest

COPY --from=builder /app/cmd/cmd /usr/local/bin/k8s-scheduler-extender

RUN adduser nonroot --disabled-password

USER nonroot

CMD ["k8s-scheduler-extender"]