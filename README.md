# go-simple

The **Go-Simple** project is a dynamic, open-source **Golang** 💻 repository that serves as a modern **blueprint** for building **scalable, clean, and production-ready web services** 🚀. It showcases best practices in architecture by strictly adhering to **Clean Architecture** 🏗️ and **domain-centric modularity**, organizing code into distinct layers (controllers, services, and repositories) to ensure the application is highly **maintainable** ✅ and easily **testable** 🧪 as it grows.

---

### ✨ Key Architectural Features

* **Dependency Management:** Uses **Uber FX** ✨ for robust dependency injection and graceful application lifecycle management.
* **Web Framework:** Leverages the **Gin** framework 🌐 for a high-performance HTTP layer.
* **Data Access:** Utilizes **SQLC** 💾 for type-safe interaction with the database.
* **Performance:** Integrates a **Redis-backed cache layer** 🧠 for speed and efficiency.
* **Communication:** Includes a **gRPC server** ⚡ to demonstrate high-performance microservice communication.
* **Core Module:** Features a fully functional **Product Management Module** 📦 with complete **CRUD** (Create, Read, Update, Delete) functionality.

---

### 📚 Documentation & Learning

This repository is an **evolving guide** to advanced Go development. All architectural decisions and feature implementations are thoroughly documented through a series of dedicated articles on Medium. By following this project, you can learn essential **advanced Go idioms**, master the use of **DTOs** (Data Transfer Objects) 💼, and implement professional-grade backend systems 💡.

[medium.com/list/gosimple-b350f5c3bdb6](https://mobinshaterian.medium.com/list/gosimple-b350f5c3bdb6)

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


## Build docker file

```
docker build -t go-simple:latest .
docker run -p 4000:4000 go-simple:latest

docker run -p 4000:4000 --env-file .env go-simple:latest

```