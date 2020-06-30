build: statik_landing wire build_server build_client

build_server:
	@CGO_ENABLED=0 go build -o ./build/bore-server ./cmd/bore-server

build_client:
	@CGO_ENABLED=0 go build -o ./build/bore ./cmd/bore

build_ui_landing:
	@if [ ! -d "web/bore-landing/dist" ]; then cd web/bore-landing && yarn build; fi

wire:
	@wire ./cmd/bore-server

statik_landing: build_ui_landing
	@if [ ! -r "internal/ui/landing/statik.go" ]; then statik -dest ./internal/ui -p landing -src ./web/bore-landing/dist; fi

install_dependencies:
	@go get github.com/jkuri/statik github.com/google/wire/cmd/...

clean:
	@rm -rf build/ internal/ui web/bore-landing/dist

.PHONY: wire build_server build_client build build_ui_landing statik_landing clean
