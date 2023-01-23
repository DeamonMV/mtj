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
    --uid "1001" "goapp"

ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["./main"]
