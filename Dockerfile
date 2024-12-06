ARG GO_VERSION=1.22-alpine3.18
ARG FROM_IMAGE=alpine:3.21

FROM golang:${GO_VERSION} AS builder

LABEL org.opencontainers.image.source="https://github.com/counterapi/counter-go"

WORKDIR /app

COPY ./ /app

RUN apk update && \
  apk add ca-certificates gettext git make curl unzip && \
  rm -rf /tmp/* && \
  rm -rf /var/cache/apk/* && \
  rm -rf /var/tmp/*

RUN make build-for-container

FROM ${FROM_IMAGE}

COPY --from=builder /app/dist/counter-linux /bin/counter

ENTRYPOINT ["counter"]
