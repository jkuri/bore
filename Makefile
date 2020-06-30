build: wire statik_landing build_server build_client

build_server:
	@CGO_ENABLED=0 go build -o ./build/bore-server ./cmd/bore-server

build_client:
	@CGO_ENABLED=0 go build -o ./build/bore ./cmd/bore

wire:
	@wire ./cmd/bore-server

statik_landing:
	@if [ ! -r "internal/ui/landing/statik.go" ]; then statik -dest ./internal/ui -p landing -src ./web/bore-landing/dist; fi

install_dependencies:
	@go get github.com/jkuri/statik github.com/google/wire/cmd/...

.PHONY: wire build_server build_client build
