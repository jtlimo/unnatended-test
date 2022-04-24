package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"unattended-test/src/deck/application"
	"unattended-test/src/deck/infrastructure"
	"unattended-test/src/deck/routes"
)

func main() {
	router := mux.NewRouter()
	db := infrastructure.New()
	deckUseCase := application.NewDeckUC(db)

	srv := routes.Server{
		Router:      router,
		DeckUseCase: deckUseCase,
	}

	srv.Register()

	log.Fatal(http.ListenAndServe(":3000", router))
}
