

build:
	go build -o bin/toker .

run: build
	bin/toker
