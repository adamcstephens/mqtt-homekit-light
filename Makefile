default:

build-arm64: export GOOS = linux
build-arm64: export GOARCH = arm64
build-arm64:
	go build -o mqtt-homekit-light.arm64linux

.PHONY: default build