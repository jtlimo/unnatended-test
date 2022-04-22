package server

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"unattended-test/card"
	"unattended-test/deck"
	"unattended-test/server/dto"
)

var s = Server{}

func TestCreateDeck(t *testing.T) {
	t.Run("create a new standard Deck", func(t *testing.T) {
		expectedDeck := dto.DeckDTO{
			Id:        "7dd13273-fabb-4223-9df6-9646c9473880",
			Shuffled:  false,
			Remaining: 52,
		}
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()
		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c9473880"
		}
		request := create("false")
		res := executeRequest(request)

		payload, _ := ioutil.ReadAll(res.Body)

		var jsonData dto.DeckDTO
		json.Unmarshal(payload, &jsonData)

		if assert.NotNil(t, jsonData) {
			assert.Equal(t, expectedDeck, jsonData)
		}
		assertStatus(t, res.Code, http.StatusOK)
	})

	t.Run("create a new custom Deck", func(t *testing.T) {
		expectedDeck := dto.DeckDTO{
			Id:        "7dd13273-fabb-4223-9df6-9646c9473881",
			Shuffled:  false,
			Remaining: 3,
		}
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()
		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c9473881"
		}
		request := create("false", "AS,2S,JS")
		res := executeRequest(request)

		payload, _ := ioutil.ReadAll(res.Body)

		var jsonData dto.DeckDTO
		json.Unmarshal(payload, &jsonData)

		if assert.NotNil(t, jsonData) {
			assert.Equal(t, expectedDeck, jsonData)
		}
		assertStatus(t, res.Code, http.StatusOK)
	})

	t.Run("create a shuffled Deck", func(t *testing.T) {
		expectedDeck := dto.DeckDTO{
			Id:        "7dd13273-fabb-4223-9df6-9646c9473882",
			Shuffled:  true,
			Remaining: 3,
		}
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()
		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c9473882"
		}
		request := create("true", "AS,2S,JS")
		res := executeRequest(request)

		payload, _ := ioutil.ReadAll(res.Body)

		var jsonData dto.DeckDTO
		json.Unmarshal(payload, &jsonData)

		if assert.NotNil(t, jsonData) {
			assert.Equal(t, expectedDeck, jsonData)
		}
		assertStatus(t, res.Code, http.StatusOK)
	})
}

func TestOpenDeck(t *testing.T) {
	t.Run("open an existent deck", func(t *testing.T) {
		cards, _ := card.NewCard([]string{"AS", "JD", "QH"})
		expectedDeck := dto.OpenDeckDTO{
			DeckDTO: dto.DeckDTO{
				Id:        "7dd13273-fabb-4223-9df6-9646c9473890",
				Shuffled:  false,
				Remaining: 3,
			},
			CardDTO: dto.ToCard(cards),
		}
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()
		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c9473890"
		}
		request := create("false", "AS,JD,QH")
		executeRequest(request)

		openRequest := open("7dd13273-fabb-4223-9df6-9646c9473890")

		res := executeRequest(openRequest)

		payload, _ := ioutil.ReadAll(res.Body)

		var jsonData dto.OpenDeckDTO
		json.Unmarshal(payload, &jsonData)

		if assert.NotNil(t, jsonData) {
			assert.Equal(t, expectedDeck, jsonData)
		}
		assertStatus(t, res.Code, http.StatusOK)
	})

	t.Run("returns not found when try to open a nonexistent deck", func(t *testing.T) {
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()

		openRequest := open("7dd13273-fabb-4223-9df6-9646c9473891")

		res := executeRequest(openRequest)

		assertStatus(t, res.Code, http.StatusNotFound)
	})

	t.Run("returns bad request when try to open a deck without remaining cards", func(t *testing.T) {
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()

		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c94738811"
		}
		createRequest := create("false", "AS,2S,JS")
		executeRequest(createRequest)

		drawRequest := draw("7dd13273-fabb-4223-9df6-9646c94738811", 3)
		executeRequest(drawRequest)

		openRequest := open("7dd13273-fabb-4223-9df6-9646c94738811")

		res := executeRequest(openRequest)

		assertStatus(t, res.Code, http.StatusBadRequest)
	})
}

func TestDrawDeck(t *testing.T) {
	t.Run("draw cards from a deck", func(t *testing.T) {
		expectedCards := []dto.CardDTO{
			{
				Value: "ACE",
				Suit:  "SPADES",
				Code:  "AS",
			},
			{
				Value: "JACK",
				Suit:  "DIAMONDS",
				Code:  "JD",
			},
		}
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()
		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c9473810"
		}
		request := create("false", "AS,JD,QH")
		executeRequest(request)

		drawRequest := draw("7dd13273-fabb-4223-9df6-9646c9473810", 2)
		res := executeRequest(drawRequest)

		payload, _ := ioutil.ReadAll(res.Body)

		var jsonData []dto.CardDTO
		json.Unmarshal(payload, &jsonData)

		if assert.NotNil(t, jsonData) {
			assert.Equal(t, expectedCards, jsonData)
		}
		assertStatus(t, res.Code, http.StatusOK)
	})

	t.Run("returns a not found error when draw cards from a nonexistent deck", func(t *testing.T) {
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()

		drawRequest := draw("7dd13273-fabb-4223-9df6-9646c9473810000", 2)
		res := executeRequest(drawRequest)

		assertStatus(t, res.Code, http.StatusNotFound)
	})

	t.Run("returns a bad request error when draw cards from a passed deck", func(t *testing.T) {
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()

		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c9473811"
		}
		request := create("false", "AS,JD,QH")
		executeRequest(request)

		drawRequest := draw("7dd13273-fabb-4223-9df6-9646c9473811", 4)
		res := executeRequest(drawRequest)

		assertStatus(t, res.Code, http.StatusBadRequest)
	})
}

func create(shuffle string, cards ...string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/deck", nil)
	queryParams := url.Values{
		"cards":   cards,
		"shuffle": []string{shuffle},
	}
	req.URL.RawQuery = queryParams.Encode()

	return req
}

func open(uuid string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/deck/%s", uuid), nil)

	return req
}

func draw(uuid string, count int) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/deck/%s/%d", uuid, count), nil)

	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)

	return rr
}
