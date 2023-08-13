#!/usr/bin/make -f

build:
	rm -rf ./build
	rm -rf ./dist
	mkdir -p build
	go build -o ./build/sonrd ./cmd/sonrd/main.go

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/sonrd-linux-amd64/sonrd ./cmd/sonrd/main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/sonrd-linux-arm64/sonrd ./cmd/sonrd/main.go

do-checksum-linux:
	cd build && sha256sum sonrd-linux-amd64/sonrd sonrd-linux-arm64/sonrd > sonr-checksum-linux

build-linux-with-checksum: build-linux do-checksum-linux

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/sonrd-darwin-amd64/sonrd ./cmd/sonrd/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/sonrd-darwin-arm64/sonrd ./cmd/sonrd/main.go

do-checksum-darwin:
	cd build && sha256sum sonrd-darwin-amd64/sonrd sonrd-darwin-arm64/sonrd > sonr-checksum-darwin

build-darwin-with-checksum: build-darwin do-checksum-darwin

build-all-with-checksum: build build-linux-with-checksum build-darwin-with-checksum

release: build-all-with-checksum
	mkdir -p dist
	tar -C build -czf dist/sonrd-linux-amd64.tar.gz sonrd-linux-amd64 sonr-checksum-linux
	tar -C build -czf dist/sonrd-linux-arm64.tar.gz sonrd-linux-arm64 sonr-checksum-linux
	tar -C build -czf dist/sonrd-darwin-amd64.tar.gz sonrd-darwin-amd64 sonr-checksum-darwin
	tar -C build -czf dist/sonrd-darwin-arm64.tar.gz sonrd-darwin-arm64 sonr-checksum-darwin
	cp ./LICENSE ./dist/LICENSE
	cp ./sonr.yml ./dist/sonr.yml
	cp ./scripts/localnet.sh ./dist/localnet.sh
