# syntax=docker/dockerfile:1.5
FROM --platform=${BUILDPLATFORM:-linux/arm64} tonistiigi/xx AS xx

FROM --platform=${BUILDPLATFORM:-linux/arm64} golang:1.23.2-alpine3.19 AS builder

COPY --from=xx / /

RUN apk add --no-cache \
    bash \
    curl \
    netcat-openbsd \
    git \
    openssh-client

WORKDIR /app

# Set GOPRIVATE to let go know which repos are private
ENV GOPRIVATE=github.com/func-it/*

# Configure known_hosts for SSH access to GitHub
RUN mkdir -p ~/.ssh && \
    chmod 700 ~/.ssh && \
    ssh-keyscan github.com >> ~/.ssh/known_hosts

# Copy only the go.mod and go.sum initially to leverage layer caching
COPY go.mod go.sum ./

# Use SSH mount so go mod download can access private repos
RUN --mount=type=ssh \
    git config --global url."git@github.com:".insteadOf "https://github.com/" && \
    go mod download -x && \
    go mod verify && \
    go install ./... && \
    go install golang.org/x/tools/cmd/goimports@latest

# Copy the rest of the source code
COPY . .

# Build the binary using Taskfile
RUN --mount=type=cache,id=build-cache,target=/root/.cache/go-build \
    xx-go --wrap && \
    go build --ldflags '-extldflags "-static" -w -s' -mod=readonly -o ./speech-to-text -tags netgo . && \
    xx-verify /app/speech-to-text

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