FROM alpine:3.14

COPY ./bin/sonrd /usr/local/bin/sonrd
COPY ./assets/networks/local/entrypoint.sh ./entrypoint.sh
RUN chmod +x /usr/local/bin/sonrd

