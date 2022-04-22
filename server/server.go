package server

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"unattended-test/database"
	"unattended-test/deck"
	"unattended-test/server/dto"
)

type Server struct {
	Router *mux.Router
	Db     *database.Database
}

func (s *Server) Setup() *mux.Router {
	s.Router = mux.NewRouter()
	s.createRoutes()
	s.initializeDB()

	return s.Router
}

func (s *Server) createRoutes() {
	s.Router.HandleFunc("/deck", s.createDeck).Methods("POST")
	s.Router.HandleFunc("/deck/{uuid}", s.openDeck).Methods("GET")
	s.Router.HandleFunc("/deck/{uuid}/{count}", s.draw).Methods("POST")
}

func (s *Server) initializeDB() {
	s.Db = database.Init()
}

func (s *Server) createDeck(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cards := strings.Split(params.Get("cards"), ",")
	shuffleStr := params.Get("shuffle")
	shuffle := false
	var err error

	if len(shuffleStr) > 0 {
		shuffle, err = strconv.ParseBool(shuffleStr)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := deck.NewDeck(cards, shuffle)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.Db.Insert(d)
	payload, err := json.Marshal(dto.ToDeck(d))

	if err != nil {
		fmt.Println("unable to encode json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (s *Server) openDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	d, err := s.Db.GetByDeckId(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if d.Remaining < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := json.Marshal(dto.ToOpenDeck(d))

	if err != nil {
		fmt.Println("unable to encode json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (s *Server) draw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	count, err := strconv.ParseInt(vars["count"], 0, 32)
	uuid := vars["uuid"]

	d, err := s.Db.GetByDeckId(uuid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	isDeckPassed := d.Remaining < 1 || int(count) > d.Remaining

	if isDeckPassed {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cards := d.Draw(int(count))

	payload, err := json.Marshal(dto.ToCard(cards))

	if err != nil {
		fmt.Println("unable to encode json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
