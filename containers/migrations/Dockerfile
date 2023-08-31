FROM migrate/migrate:4

COPY --from=ghcr.io/ufoscout/docker-compose-wait:2.12 /wait /wait

COPY ./entrypoint.sh /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
