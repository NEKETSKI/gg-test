FROM golang:1.20-bullseye AS builder
WORKDIR /go/src/github.com/NEKETSKY/gg-test
COPY . .
ARG APP_VERSION
RUN go build -ldflags="-X 'main.version=${APP_VERSION}'" -o .build/gg-test ./cmd/gg-test

FROM debian:11-slim as runner
WORKDIR /bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/NEKETSKY/gg-test/.build/gg-test .
RUN addgroup --system gg-test && adduser --system --group gg-test
USER gg-test
ENTRYPOINT ["gg-test"]
