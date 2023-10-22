#!/usr/bin/env bash

set -eu -o pipefail

cd "$(dirname $0)/../.."

docker container run -it --rm \
  -u $(id -u):$(id -g) \
  -v $(pwd):/app \
  -w /app \
  openapitools/openapi-generator-cli:v7.0.1 \
  generate -i /app/docs/api/openapi.yml -g typescript-fetch -o /app/realworld-web/app/lib/oapi/generated/
