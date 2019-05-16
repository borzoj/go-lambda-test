wire:
	wire github.com/borzoj/go-lambda-test/cmd

build: clean wire
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler ./cmd

.PHONY: clean
clean:
	rm -rf ./bin

run: build
	go run ./cmd

.PHONY: deploy
deploy: build
	serverless deploy --verbose
