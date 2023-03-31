FROM golang:1.19.5-bullseye AS build-env

LABEL org.opencontainers.image.source https://github.com/sonr-io/sonr
LABEL org.opencontainers.image.description Sonr Blockchain Node Daemon as Container

WORKDIR /go/src/github.com/sonr-io/sonr

RUN apt-get update -y
RUN apt-get install git -y

COPY . .

RUN make build

FROM golang:1.19.5-bullseye

RUN apt-get update -y
RUN apt-get install ca-certificates jq -y

WORKDIR /root

COPY --from=build-env /go/src/github.com/sonr-io/sonr/build/sonrd /usr/bin/sonrd

EXPOSE 26656 26657 1317 9090 8545 8546

CMD ["sonrd"]
