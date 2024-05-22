package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/ret0rn/password-generator/pkg/routes"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	routes.RegisterPassGenRouter(router)
	port := ":80"
	logrus.Infof("Server start at port %s", port)
	server := &http.Server{
		Addr:              port,
		ReadHeaderTimeout: 3 * time.Second,
		Handler:           router,
	}
	logrus.Fatal(server.ListenAndServe())
}
