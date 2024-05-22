package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/ret0rn/password-generator/pkg/routes"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := ":80"
	routes.RegisterPassGenRouter(router)
	logrus.Infof("Server start at port %s", port)
	logrus.Fatal(http.ListenAndServe(port, router))
}
