FROM alpine:latest

RUN apk --update --no-cache add curl
LABEL org.opencontainers.image.title="Sonr"
LABEL org.opencontainers.image.description="An IBC AppChain for Decentralized Identity"
LABEL org.opencontainers.image.source="https://github.com/sonrhq/sonr"
LABEL org.opencontainers.image.licenses="OpenGL 3.0"
LABEL org.opencontainers.image.documentation="https://sonr.guide"
LABEL org.opencontainers.image.vendor="Sonr Inc."

COPY ./bin/create-account /usr/bin/create-account
COPY ./bin/dendrite /usr/bin/dendrite
COPY ./bin/generate-config /usr/bin/generate-config
COPY ./bin/generate-keys /usr/bin/generate-keys
COPY ./bin/hway /usr/bin/hway
COPY ./bin/sonrd /usr/bin/sonrd

ENTRYPOINT ["/usr/bin/sonrd"]
