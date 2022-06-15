package main

import (
	"log"
	"net/http"

	"github.com/pedroreiscoder/product-rest-api/data"
	"github.com/pedroreiscoder/product-rest-api/router"
)

func main() {
	data.Init()
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3200", r))
}
