
ARG GO_VERSION="1.19"
ARG RUNNER_IMAGE="debian:bullseye-slim"

FROM ${RUNNER_IMAGE}

LABEL org.opencontainers.image.source https://github.com/sonrhq/sonr
LABEL org.opencontainers.image.description "Sonr Validator node container"

# Copy sonrd binary and config
COPY bin/sonrd /usr/local/bin/sonrd
COPY scripts/entrypoint.sh entrypoint.sh

# Expose ports
EXPOSE 26657
EXPOSE 1317
EXPOSE 26656
EXPOSE 8080
