ARG RUNNER_IMAGE="debian:bullseye-slim"

FROM ${RUNNER_IMAGE}

LABEL org.opencontainers.image.source https://github.com/sonrhq/sonr
LABEL org.opencontainers.image.description "Sonr Validator node container"

# Copy sonrd binary and config
COPY gopath/bin/sonrd /usr/local/bin/sonrd

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
