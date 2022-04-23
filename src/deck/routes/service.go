package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
	"unattended-test/src/deck/application"
	"unattended-test/src/deck/domain"
	"unattended-test/src/deck/routes/dto"
)

type Server struct {
	Router      *mux.Router
	DeckUseCase *application.DeckUC
}

func (s *Server) Register() {
	s.Router.HandleFunc("/deck", s.createDeck).Methods("POST")
	s.Router.HandleFunc("/deck/{uuid}", s.openDeck).Methods("GET")
	s.Router.HandleFunc("/deck/{uuid}/{count}", s.draw).Methods("POST")
}

func (s *Server) createDeck(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	cards := strings.Split(params.Get("cards"), ",")
	shuffleStr := params.Get("shuffle")
	shuffle := false
	cards = sanitizeParams(cards)
	var err error

	if len(shuffleStr) > 0 {
		shuffle, err = strconv.ParseBool(shuffleStr)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	d, err := domain.New(cards, shuffle)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	s.DeckUseCase.Create(d)

	payload, err := json.Marshal(dto.ToDeck(d))

	if err != nil {
		fmt.Println("Unable to encode json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (s *Server) openDeck(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	d, err := s.DeckUseCase.Open(uuid)
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
		fmt.Println("Unable to encode json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (s *Server) draw(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	count, err := strconv.ParseInt(vars["count"], 0, 32)
	uuid := vars["uuid"]

	cards, err := s.DeckUseCase.Draw(int(count), uuid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := json.Marshal(dto.ToCard(cards))

	if err != nil {
		fmt.Println("unable to encode json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func sanitizeParams(cards []string) []string {
	var params []string
	for _, card := range cards {
		params = append(params, strings.ToUpper(card))
	}

	return params
}
