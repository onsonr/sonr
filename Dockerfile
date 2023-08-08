# ! ||--------------------------------------------------------------------------------||
# ! ||                                  Sonrd Builder                                 ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux golang:1.19-alpine AS builder

ENV SONR_VERSION=master
RUN apk add --update gcc g++ make git curl

WORKDIR /root
RUN git clone --depth 1 --branch ${SONR_VERSION} https://github.com/sonrhq/core.git sonr

WORKDIR /root/sonr
RUN go build -o ./build/sonrd ./cmd/sonrd/main.go


# ! ||--------------------------------------------------------------------------------||
# ! ||                               Sonr in Production                               ||
# ! ||--------------------------------------------------------------------------------||
FROM --platform=linux alpine

# Copy sonrd binary and config
COPY --from=builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.docker.yml sonr.yml
COPY scripts scripts
ENV SONR_LAUNCH_CONFIG=/sonr.yml

# Copy IceFire binaries and config
COPY build/bin/IceFireDB /usr/local/bin/icefirekv
COPY build/bin/IceFireDB-SQLite /usr/local/bin/icefiresql
COPY build/config/config.sql.yaml config.sql.yaml
COPY build/db/read.txt /db/read.txt

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
EXPOSE 9090
