# syntax=docker/dockerfile:1.0.0-experimental
FROM golang:1.14-alpine as builder

# Default to using go modules and go proxy
ENV GO111MODULE on
ENV GOPRIVATE github.com/getoutreach/*
ENV GOPROXY https://proxy.golang.org

WORKDIR /src

# Ensure that we're using the latest git, openssh, and ca-certs
RUN apk add --no-cache --update \
  openssh-client \
  git \
  ca-certificates

# Add go.mod and go.sum first to maximize caching
COPY ./go.mod ./go.sum ./
RUN --mount=type=ssh set -e -x; \
    mkdir -p ~/.ssh; \
    printf "Host github.com\n  HostName github.com\n  User git" > ~/.ssh/config; \
    ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts; \
    git config --global url."git@github.com:".insteadOf "https://github.com/"; \
    go mod download

COPY . .

RUN go build -o /src/bin/serverd ./cmd/serverd/...

FROM alpine:latest

COPY --from=builder /src/bin/serverd /usr/local/bin/serverd