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

## How to Use
`GET /api/products` Returns a list of products  

**Response example:**
```
[
    {
        "id": 1,
        "name": "Cup",
        "price": 5.35
    },
    {
        "id": 2,
        "name": "Baseball Hat",
        "price": 42.54
    }
]
```

`GET /api/product/{id}` Returns the product with the specified id  

**Response example:**
```
{
    "id": 1,
    "name": "Cup",
    "price": 5.35
}
```
