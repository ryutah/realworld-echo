#!/usr/bin/env sh

set -ue

/wait

go run /app/tools/testdata-loader/main.go \
  -p /app/resources/testdata \
  -c ${CONNECTION_NAME}
