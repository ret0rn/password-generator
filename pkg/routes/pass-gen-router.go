package routes

import (
	"github.com/gorilla/mux"

	"github.com/ret0rn/password-generator/pkg/controllers"
)

var RegisterPassGenRouter = func(router *mux.Router) {
	router.HandleFunc("/generate/", controllers.GeneratePassword).Methods("POST")
}
