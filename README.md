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

**Response:**
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

**Response:**
```
{
    "id": 1,
    "name": "Cup",
    "price": 5.35
}
```

`POST /api/product` Creates a new product and returns the created product

**Request body:**
```
{
    "name": "Shoe",
    "price": 25.99
}
```

**Response:**
```
{
    "id": 3,
    "name": "Shoe",
    "price": 25.99
}
```

`PUT /api/product/{id}` Updates an existing product

**Request body:**
```
{
    "name": "New Cup",
    "price": 10.70
}
```

**Response:**
```
204 No Content
```
