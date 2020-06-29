build: wire build_server

build_server:
	@CGO_ENABLED=0 go build -o ./build/bore-server ./cmd/bore-server

wire:
	@wire ./cmd/...

install_dependencies:
	@go get github.com/google/wire/cmd/...

.PHONY: wire build_server build
