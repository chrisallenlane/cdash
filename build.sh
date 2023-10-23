#!/bin/sh

# compile AMD64 for Linux, OSX, and Windows
env GOOS=darwin GOARCH=amd64 go build  -o /tmp/cdash-darwin-amd64  .
env GOOS=darwin GOARCH=arm64 go build  -o /tmp/cdash-darwin-arm64  .
env GOOS=linux GOARCH=amd64 go build   -o /tmp/cdash-linux-amd64   .
env GOOS=windows GOARCH=amd64 go build -o /tmp/cdash-win-amd64.exe .
