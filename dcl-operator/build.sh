#!/bin/sh

VERSION=$(git describe --tags --abbrev=0)

go build -ldflags "-X main.VERSION=$VERSION" .
