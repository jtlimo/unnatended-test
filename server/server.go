package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"unattended-test/database"
)

type Server struct {
	Router *mux.Router
	Db     *database.Database
}

func (s *Server) Setup() {
	s.Router = mux.NewRouter()
	s.createRoutes()
	s.initializeDB()

	log.Fatal(http.ListenAndServe(":8080", s.Router))
}

func (s *Server) createRoutes() {
	s.Router.HandleFunc("/deck", createDeck).Methods("POST").Queries("cards", "{cards}")
	http.Handle("/", s.Router)
}

func (s *Server) initializeDB() {
	s.Db = database.Init()
}

func createDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deck created: %v\n", vars["cards"])
}
