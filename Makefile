build: wire build_server build_client

build_server:
	@CGO_ENABLED=0 go build -o ./build/bore-server ./cmd/bore-server

build_client:
	@CGO_ENABLED=0 go build -o ./build/bore ./cmd/bore

wire:
	@wire ./cmd/bore-server

install_dependencies:
	@go get github.com/google/wire/cmd/...

.PHONY: wire build_server build_client build
