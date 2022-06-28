# product-rest-api
REST API developed using Go, PostgreSQL and Gorilla Mux that allows a user to create, read, update and delete products from a store

## Requirements
Go v1.18+  
PostgreSQL v14.4+

## Installation
After cloning the repository, in the root directory run the following command:
```go
go get .
```

Go to the `db.go` file inside the data folder and replace the connection information with the settings of your Postgres instance  

In the root directory run the command:
```go
go run main.go
```
