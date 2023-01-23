FROM alpine:3.17

COPY docker-entrypoint.sh /docker-entrypoint.sh
COPY main /main

RUN apk add --no-cache libc6-compat su-exec bash gettext \
 &&  chmod +x /docker-entrypoint.sh \
 && adduser \
    --disabled-password \
    --gecos "" \
    --home "/var/www/go" \
    --no-create-home \
    --uid "1001" "www-data"

ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["./main"]
>>>>>>> 4465a16 (update dockerfile, add entrypoint)
