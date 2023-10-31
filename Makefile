build:
	@go build -o bin/go-rest-api

run: build
	@./bin/go-rest-api

test:
	@go test -v ./...
