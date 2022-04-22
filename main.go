package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"unattended-test/database"
	"unattended-test/server"
)

func main() {
	router := mux.NewRouter()
	db := database.New()

	srv := server.Server{
		Router: router,
		Db:     db,
	}
	srv.CreateRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
