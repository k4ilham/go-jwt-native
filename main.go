package main

import (
	"go-jwt/configs"
	"go-jwt/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	log.Println("Server running on port 8060")
	http.ListenAndServe(":8060", router)
}
