package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"unattended-test/deck"
)

var expectedDeck = deck.Deck{
	Id:        "7dd13273-fabb-4223-9df6-9646c9473880",
	Shuffled:  false,
	Remaining: 52,
}
var s = Server{}

func TestCreateDeck(t *testing.T) {
	t.Run("create a new standard Deck", func(t *testing.T) {
		router := s.Setup()
		localServer := httptest.NewServer(router)
		defer localServer.Close()

		var old = deck.GenerateNewUUID
		defer func() { deck.GenerateNewUUID = old }()
		deck.GenerateNewUUID = func() string {
			return "7dd13273-fabb-4223-9df6-9646c9473880"
		}
		request := create()
		res := executeRequest(request)

		d := s.Db.Get()
		if assert.NotNil(t, d) {
			assert.Equal(t, expectedDeck, d)
		}
		assertStatus(t, res.Code, 200)
	})
}

func create(cards ...string) *http.Request {
	var req, _ *http.Request
	if len(cards) > 0 {
		req, _ = http.NewRequest(http.MethodPost, fmt.Sprintf("/deck?cards=%s", cards), nil)
	} else {
		req, _ = http.NewRequest(http.MethodPost, "/deck", nil)

	}
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
