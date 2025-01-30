

build:
	go build -o bin/toker cmd/main.go

run: build
	bin/toker
