

build:
	go build -o bin/toker cmd/main.go

run: build
	bin/toker

wire:
	wire ./cmd/boot

install-deps:
	go install github.com/google/wire/cmd/wire@latest
