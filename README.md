# go + gin + serverless + lambda


Simple REST API which can be run as AWS Lambda or as local HTTP server.

Uses:
1. [gin framework](https://github.com/gin-gonic/gin) for request handling
2. [apex gateway](https://github.com/apex/gateway) to serve the same code as Lambda and local HTTP
3. [serverless framework](https://serverless.com/) for nicer Lambda definitions and extra functionality
4. [wire](https://github.com/google/wire) for dependency injection


## Requirements:

```
make
serverless
go 1.12
```

## Run locally

```
make run
```

The server will run on http://localhost:3000/

## Deploy to AWS
```
make deploy
```
If deployment is successful serverless will output the Lambda root URL

## API

### GET /foo
Sample GET request

### GET /weather/{city}

Sample GET request which makes a request to another API endpoint

### POST /ask
Sample POST request. Post body:

```
{ "body": "question body" }
```


## Further development

1. Dockerize to simplify setup of dependencies
1. Add tests
1. Mocking library for tests
1. Logging library for structured logs