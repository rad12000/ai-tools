#!/usr/bin/env bash
set -euo pipefail

for target in darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64 windows/arm64; do
  os=${target%/*} arch=${target#*/}
  ext=$([ "$os" = windows ] && echo .exe || true)
  CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -ldflags="-s -w" -trimpath \
    -o "dist/schema-template_${os}_${arch}${ext}" main.go
done
