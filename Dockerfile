FROM golang:1.14-alpine as builder

ENV USER=appuser
ENV UID=10001

ENV CGO_ENABLED=0

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

ARG VERSION=v0

WORKDIR /app

COPY . /app

RUN apk update && apk add --no-cache make ca-certificates \
    && update-ca-certificates \
    && BIN_VERSION=$VERSION make build

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /usr/bin
COPY --from=builder /app/ozone /usr/bin/ozone

USER appuser:appuser

WORKDIR /app

EXPOSE 8080

ENTRYPOINT ["/usr/bin/ozone", "-r", "/app", "-n", "0.0.0.0"]

VOLUME ["/app"]
