build:
	go build -o bin/api -v

run: build
	./bin/api

test:
	go test -v go-reference-structure/tests