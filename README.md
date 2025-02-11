# Golang Microservices Assignment

## Overview
This microservice reads a large `ports.json` file and updates port records in an **in-memory database**. The project follows **clean architecture, Golang best practices, and JSON streaming** to handle large files.

## How to Run
```sh
git clone https://github.com/chetanji028/golang-microservices-assignment.git
cd golang-microservices-assignment
go mod tidy
go run cmd/main.go
