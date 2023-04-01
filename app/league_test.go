package poker_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/igorakimy/poker"
)

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.NewStubPlayerStore(nil, nil, wantedLeague)
		server := mustMakePlayerServer(t, store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromRequest(t, response.Body)
		poker.AssertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		poker.AssertContentType(t, response, "application/json")
	})
}

func getLeagueFromRequest(t *testing.T, body io.Reader) (league []poker.Player) {
	t.Helper()
	if err := json.NewDecoder(body).Decode(&league); err != nil {
		t.Fatalf(
			"Unable to parse response from server %q into slice of Player, '%v'",
			body,
			err,
		)
	}
	return
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func assertLeague(t *testing.T, got, want []poker.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
