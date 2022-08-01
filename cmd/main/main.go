package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ret0rn/password-generator/pkg/routes"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":80"
	routes.RegisterPassGenRouter(router)
	log.Printf("Server start at port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
