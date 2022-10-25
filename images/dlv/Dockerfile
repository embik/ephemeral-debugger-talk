FROM golang:1.19.8 as builder

WORKDIR /go

RUN git clone https://github.com/go-delve/delve && cd delve && git checkout v1.20.1 && CGO_ENABLED=0 go build -o dlv ./cmd/dlv

FROM alpine:3.16

COPY --from=builder /go/delve/dlv /bin/dlv

ENTRYPOINT ["/bin/sh"]
