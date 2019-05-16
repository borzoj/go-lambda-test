build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler ./cmd

.PHONY: clean
clean:
	rm -rf ./bin

run:
	go run ./cmd

.PHONY: deploy
deploy: clean build
	serverless deploy --verbose
