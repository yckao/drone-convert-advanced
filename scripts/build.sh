#!/bin/sh
set -e
set -x

# linux
GOOS=linux GOARCH=amd64 go build -o ./release/linux/amd64/drone-convert-advanced ./cmd/drone-convert-advanced
GOOS=linux GOARCH=arm64 go build -o ./release/linux/arm64/drone-convert-advanced ./cmd/drone-convert-advanced
GOOS=linux GOARCH=arm   go build -o ./release/linux/arm/drone-convert-advanced ./cmd/drone-convert-advanced