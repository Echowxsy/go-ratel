BINARY = ratel
VERSION = 1.0.0
GOOS = darwin
GOARCH = amd64
GOOSLINUX = linux
GOOSWINDOWS = windows

build:
	@export GOOS=${GOOS}; \
  	export GOARCH=${GOARCH}; \
  	go build -o ratel-${GOOS}-${GOARCH}-${VERSION}-cn

build-linux:
	@export GOOS=${GOOSLINUX}; \
	export GOARCH=${GOARCH}; \
	go build -o ratel-${GOOSLINUX}-${GOARCH}-${VERSION}-cn

build-windows:
	@export GOOS=${GOOSWINDOWS}; \
	export GOARCH=${GOARCH}; \
	go build -o ratel-${GOOSWINDOWS}-${GOARCH}-${VERSION}-cn.exe


test: build
	@./$(BINARY) -h 39.105.65.8 -p 1024

all: build build-linux build-windows

.PHONY: build test
