# go-simple

This is a simple Go project structure with Gin framework and Swagger documentation.

## Run project

go run cmd/server/main.go

## Swagger address

http://127.0.0.1:4000/swagger/index.html#/

## Generate Swagger Documentation

swag init --parseDependency --parseInternal -g cmd/server/main.go


## Run docker compose

docker compose up -d

## SQLC Generator

```
sqlc generate
```

## Run tests

```
go test ./test/server_test.go

```

## Run Sonar

```
docker run --name sonarqube \
  -p 9000:9000 \
  sonarqube:latest


go test -coverpkg=./... -coverprofile=coverage.out ./test

go test ./... -json > report.json

go tool cover -func=coverage.out
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



## Run project

```
docker compose up
go run cmd/server/main.go
```


## Generate Proto

```
protoc -I.   --go_out=. --go_opt=paths=source_relative   --go-grpc_out=. --go-grpc_opt=paths=source_relative   --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative   api/proto/product/v1/product.proto
```


## Install GRPC-Gateway

```
go get google.golang.org/grpc
go get -tool github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get -tool github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go get -tool google.golang.org/protobuf/cmd/protoc-gen-go
go get -tool google.golang.org/grpc/cmd/protoc-gen-go-grpc
```


## deploy on Liara

```
npm install -g @liara/cli
liara login
liara deploy
```