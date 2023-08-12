#!/usr/bin/make -f

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/sonrd-linux-amd64/sonrd ./cmd/sonrd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/sonrd-linux-arm64/sonrd ./cmd/sonrd/main.go

do-checksum-linux:
	cd build && sha256sum sonrd-linux-amd64/sonrd sonrd-linux-arm64/sonrd > sonr-checksum-linux

build-linux-with-checksum: build-linux do-checksum-linux

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/sonrd-darwin-amd64/sonrd ./cmd/sonrd/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/sonrd-darwin-arm64/sonrd ./cmd/sonrd/main.go

build-all: build-linux build-darwin

do-checksum-darwin:
	cd build && sha256sum sonrd-darwin-amd64 sonrd-darwin-arm64 > sonr-checksum-darwin

build-darwin-with-checksum: build-darwin do-checksum-darwin

build-with-checksum: build-linux-with-checksum build-darwin-with-checksum
