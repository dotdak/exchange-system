#!/bin/bash
BINARY=app
VERSION=1.0.0
BUILD_TIME=`date +%FT%T%z`
BUILD_GO_VERSION=`go version`

# Setup the -ldflags option for build info here, interpolate the variable values
# notice: replace the path with your versionInfo module path
LDFLAGS="-w -s \
-X 'version.BinVersion=${VERSION}' \
-X 'version.BuildTime=${BUILD_TIME}' \
-X 'version.BuildGoVersion=${BUILD_GO_VERSION}' \
"

go build -ldflags "${LDFLAGS}" -o ${BINARY}