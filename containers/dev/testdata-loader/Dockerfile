FROM golang:1.21.0-bookworm

COPY --from=ghcr.io/ufoscout/docker-compose-wait:2.12 /wait /wait

COPY ./entrypoint.sh /entrypoint.sh

WORKDIR /app

ENTRYPOINT ["/entrypoint.sh"]
