build:
	@go build -o bin/lbs

run: build
	@./bin/lbs

test:
	@go test -v ./...
