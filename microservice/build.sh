#!/bin/bash

# Create output directory
mkdir -p builds

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o builds/zwartn-microserviceA.exe main.go

# Build for arm macOS
GOOS=darwin GOARCH=arm64 go build -o builds/zwartn-microserviceA-mac-arm64 main.go

# Create checksums
cd builds
sha256sum * > checksums.txt
cd ..