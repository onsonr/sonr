ARG BASE_VERSION
FROM golang:${BASE_VERSION} AS init-env

RUN apk add --update --no-cache curl make git libc-dev bash gcc linux-headers eudev-dev ncurses-dev

ARG TARGETARCH
ARG BUILDARCH
ARG GITHUB_ORGANIZATION
ARG REPO_HOST
ARG GITHUB_REPO
ARG WASMVM_VERSION

WORKDIR /go/src/${REPO_HOST}/${GITHUB_ORGANIZATION}/${GITHUB_REPO}

# Download CosmWasm libwasmvm if found
RUN set -eux; \
    export ARCH=$(uname -m); \
    if [ ! -z "${WASMVM_VERSION}" ]; then\
      WASMVM_REPO=$(echo $WASMVM_VERSION | awk '{print $1}');\
      WASMVM_VERS=$(echo $WASMVM_VERSION | awk '{print $2}');\
      wget -O /lib/libwasmvm_muslc.a https://${WASMVM_REPO}/releases/download/${WASMVM_VERS}/libwasmvm_muslc.$(uname -m).a;\
    fi;

ARG BUILD_DIR

ADD ${BUILD_DIR}/go.mod ${BUILD_DIR}/go.sum ./

# Download go mod dependencies, if there is no custom build directory
# Note: a custom build dir indicates a monorepo with potential dependencies we can't anticipate atm
RUN set -eux; \
    if [[ "${BUILD_DIR}" == "." ]]; then \
      go mod download; \
    fi;

# Use minimal busybox from infra-toolkit image for final scratch image
FROM ghcr.io/strangelove-ventures/infra-toolkit:v0.1.4 AS infra-toolkit
RUN addgroup --gid 1025 -S heighliner && adduser --uid 1025 -S heighliner -G heighliner

# Use ln and rm from full featured busybox for assembling final image
FROM busybox:1.34.1-musl AS busybox-full

# Build part 1 of the final image
FROM scratch AS final-part1

LABEL org.opencontainers.image.source="https://github.com/strangelove-ventures/heighliner"

WORKDIR /bin

# Install ln (for making hard links) and rm (for cleanup) from full busybox image (will be deleted, only needed for image assembly)
COPY --from=busybox-full /bin/ln /bin/rm ./

# Install minimal busybox image as shell binary (will create hardlinks for the rest of the binaries to this data)
COPY --from=infra-toolkit /busybox/busybox /bin/sh

# Install jq
COPY --from=infra-toolkit /usr/local/bin/jq /bin/

# Add hard links for read-only utils
# Will then only have one copy of the busybox minimal binary file with all utils pointing to the same underlying inode
RUN for b in \
  cat \
  date \
  df \
  du \
  env \
  grep \
  head \
  less \
  ls \
  md5sum \
  pwd \
  sha1sum \
  sha256sum \
  sha3sum \
  sha512sum \
  sleep \
  stty \
  tail \
  tar \
  tee \
  tr \
  watch \
  which \
  ; do ln sh $b; done

# Remove write utils
RUN rm ln rm

# Install trusted CA certificates
COPY --from=infra-toolkit /etc/ssl/cert.pem /etc/ssl/cert.pem

# Install heighliner user
COPY --from=infra-toolkit /etc/passwd /etc/passwd
COPY --from=infra-toolkit --chown=1025:1025 /home/heighliner /home/heighliner
COPY --from=infra-toolkit --chown=1025:1025 /tmp /tmp

# Install chain binary
FROM init-env AS build-env

ARG BUILD_TARGET
ARG BUILD_ENV
ARG BUILD_TAGS
ARG PRE_BUILD
ARG BUILD_DIR

# This Dockerfile  is the same as native.Dockerfile except that the chain code is sourced from the
# current working directory instead of a remote git repository.
ADD . .

RUN set -eux; \
    export CGO_ENABLED=1 LDFLAGS='-linkmode external -extldflags "-static"'; \
    if [ ! -z "$PRE_BUILD" ]; then sh -c "${PRE_BUILD}"; fi; \
    if [ ! -z "$BUILD_TARGET" ]; then \
      if [ ! -z "$BUILD_ENV" ]; then export ${BUILD_ENV}; fi; \
      if [ ! -z "$BUILD_TAGS" ]; then export "${BUILD_TAGS}"; fi; \
      if [ ! -z "$BUILD_DIR" ]; then cd "${BUILD_DIR}"; fi; \
      sh -c "${BUILD_TARGET}"; \
    fi

# Copy all binaries to /root/bin, for a single place to copy into final image.
# If a colon (:) delimiter is present, binary will be renamed to the text after the delimiter.
RUN mkdir /root/bin
ARG RACE
ARG BINARIES
ENV BINARIES_ENV ${BINARIES}
RUN bash -c 'set -eux;\
  BINARIES_ARR=();\
  IFS=, read -ra BINARIES_ARR <<< "$BINARIES_ENV";\
  for BINARY in "${BINARIES_ARR[@]}"; do\
    BINSPLIT=();\
    IFS=: read -ra BINSPLIT <<< "$BINARY";\
    BINPATH=${BINSPLIT[1]+"${BINSPLIT[1]}"};\
    BIN="$(eval "echo "${BINSPLIT[0]+"${BINSPLIT[0]}"}"")";\
    if [ ! -z "$RACE" ] && GOVERSIONOUT=$(go version -m $BIN); then\
      if echo $GOVERSIONOUT | grep build | grep "-race=true"; then\
        echo "Race detection is enabled in binary";\
      else\
        echo "Race detection not enabled in binary!";\
        exit 1;\
      fi;\
    fi;\
    if [ ! -z "$BINPATH" ]; then\
      if [[ $BINPATH == *"/"* ]]; then\
        mkdir -p "$(dirname "${BINPATH}")";\
        cp "$BIN" "${BINPATH}";\
      else\
        cp "$BIN" "/root/bin/${BINPATH}";\
      fi;\
    else\
      cp "$BIN" /root/bin/;\
    fi;\
  done'

RUN mkdir -p /root/lib
ARG LIBRARIES
ENV LIBRARIES_ENV ${LIBRARIES}
RUN bash -c 'set -eux;\
  LIBRARIES_ARR=($LIBRARIES_ENV); for LIBRARY in "${LIBRARIES_ARR[@]}"; do cp $LIBRARY /root/lib/; done'

# Move final binary to the final image
FROM final-part1 as final
WORKDIR /bin

# Install chain binaries
COPY --from=build-env /root/bin /bin

# Install libraries
COPY --from=build-env /root/lib /lib

WORKDIR /home/heighliner
USER heighliner
