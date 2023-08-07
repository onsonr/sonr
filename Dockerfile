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

# Copy the sonrd binary from the builder stage and local config
COPY --from=builder /root/sonr/build/sonrd /usr/local/bin/sonrd
COPY sonr.yml .
COPY scripts scripts

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
EXPOSE 9090
