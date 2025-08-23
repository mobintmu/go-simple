# go-simple

This is a simple Go project structure with Gin framework and Swagger documentation.

## Run project

go run cmd/server/main.go

## Generate Swagger Documentation

swag init --parseDependency --parseInternal -g cmd/server/main.go

