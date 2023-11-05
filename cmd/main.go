package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	app := mux.NewRouter()

	dep := InitDependencies()

	SetupRoutes(app, dep)

	fmt.Println("Server listening at :8080 ")
	log.Fatal(http.ListenAndServe(":8080", app))
}
