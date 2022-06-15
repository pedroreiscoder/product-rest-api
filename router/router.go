package router

import (
	"github.com/gorilla/mux"
	"github.com/pedroreiscoder/product-rest-api/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/api/product/{id}", controllers.GetProduct).Methods("GET")

	return router
}
