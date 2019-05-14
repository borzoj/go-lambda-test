build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/handler cmd/main.go

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: clean build
	serverless deploy --verbose
