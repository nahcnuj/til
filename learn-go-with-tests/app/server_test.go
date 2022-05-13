package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewServer(&store)

	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestScoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewServer(store)

	t.Run("record wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		AssertPlayerWin(t, store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("return the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewServer(&store)

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, http.StatusOK, response.Code)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response.Result(), jsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := NewServer(&StubPlayerStore{})

		request := newGetGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}

func newGetScoreRequest(name string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
}

func newGetLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}

func newGetGameRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/game", nil)
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
	t.Helper()
	if err := json.NewDecoder(body).Decode(&league); err != nil {
		t.Fatalf("unable to parse response %q into slice of Player, '%v'", body, err)
	}
	return
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("status code is wrong, got %d, want %d", got, want)
	}
}

func assertContentType(t testing.TB, response *http.Response, want string) {
	got := response.Header.Get("content-type")
	if got != want {
		t.Errorf("content-type is wrong, got %q, want %q", got, want)
	}
}

func assertLeague(t testing.TB, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("league table is wrong, got %v, want %v", got, want)
	}
}
