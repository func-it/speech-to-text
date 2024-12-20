# syntax=docker/dockerfile:1

FROM --platform=${BUILDPLATFORM:-linux/arm64} tonistiigi/xx AS xx

FROM --platform=${BUILDPLATFORM:-linux/arm64} golang:1.23.2-alpine3.19 AS builder

COPY --from=xx / /

RUN apk add --no-cache \
    bash \
    curl \
    ffmpeg \
    netcat-openbsd \
    git

# Install Taskfile
RUN sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d v3.40.0

WORKDIR /app

COPY ./go.mod ./go.sum ./

ARG TARGETPLATFORM
ARG CGO_ENABLED=0

RUN --mount=type=cache,id=build-cache,target=/root/.cache/go-build \
    go mod download -x \
    && go mod verify \
    && go install ./... \
    && go install golang.org/x/tools/cmd/goimports@latest

COPY / .
RUN --mount=type=cache,id=build-cache,target=/root/.cache/go-build \
    xx-go --wrap \
    && task speech-to-text-build \
    && xx-verify /app/speech-to-text

FROM alpine:3.19
RUN apk add --no-cache \
    ffmpeg \
    netcat-openbsd \
    bash

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/speech-to-text .

LABEL org.opencontainers.image.source="https://github.com/func-it/speech-to-text"

CMD ["./speech-to-text"]
