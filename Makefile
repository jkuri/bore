build: wire build_server

build_server:
	@go build -o ./build/bore-server ./cmd/bore-server

wire:
	@wire ./cmd/...

.PHONY: wire build_server build
