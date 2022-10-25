FROM golang:1.19.8 as builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o app .

FROM alpine:3.16
LABEL maintainer="marvin@kubermatic.com"

COPY --from=builder /build/app /bin/app

ENTRYPOINT ["/bin/app"]
