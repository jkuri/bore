build: wire build_server

build_server:
	@go build -o ./build/bored ./cmd/bored

wire:
	@wire ./cmd/...

.PHONY: wire build_server build
