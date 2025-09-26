# go-simple

This is a simple Go project structure with Gin framework and Swagger documentation.

## Run project

go run cmd/server/main.go

## Swagger address

http://127.0.0.1:4000/swagger/index.html#/

## Generate Swagger Documentation

swag init --parseDependency --parseInternal -g cmd/server/main.go


## Run tests

```
go test ./test/server_test.go

```

## Run Sonar

```
docker run --name sonarqube \
  -p 9000:9000 \
  sonarqube:latest


go test -coverprofile=coverage.out ./...

go test ./... -json > report.json
```


## Add sonar test

```
sonar-scanner \
  -Dsonar.projectKey=go-simple \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://127.0.0.1:9000 \
  -Dsonar.token=sqp_XXXX
```

```
export SONAR_HOST_URL=http://your-sonarqube-server.com
export SONAR_TOKEN=your-sonar-token-here
```
