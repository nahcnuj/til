package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
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
	server, err := NewServer(&store)
	if err != nil {
		t.Fatal("could not start a server")
	}

	t.Run("return Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertStatus(t, response, http.StatusOK)
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertStatus(t, response, http.StatusOK)
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusNotFound)
	})
}

func TestScoreWins(t *testing.T) {
	store := &StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := mustMakePlayerServer(t, store)

	t.Run("record wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusAccepted)

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

		store := &StubPlayerStore{nil, nil, wantedLeague}
		server := mustMakePlayerServer(t, store)

		request := newGetLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertStatus(t, response, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response.Result(), jsonContentType)
	})
}

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &StubPlayerStore{})

		request := newGetGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response, http.StatusOK)
	})

	t.Run("start a game with 3 players and declare Ruth as the winner", func(t *testing.T) {
		game := &SpyGame{}
		winner := "Ruth"
		store := &StubPlayerStore{}
		server := httptest.NewServer(mustMakePlayerServer(t, store, game))
		defer server.Close()

		wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"
		ws := mustDialWebSocket(t, wsURL)
		defer ws.Close()

		writeMessageToWebSocket(t, ws, "3")
		writeMessageToWebSocket(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		AssertGameStartedWith(t, game, 3)
		AssertGameFinishedWith(t, game, winner)
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

func mustMakePlayerServer(t testing.TB, store PlayerStore) *PlayerServer {
	t.Helper()
	server, err := NewServer(store)
	if err != nil {
		t.Fatal("could not start a server")
	}
	return server
}

func mustDialWebSocket(t testing.TB, url string) *websocket.Conn {
	t.Helper()
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("could not open a websocket connection on %s: %v", url, err)
	}
	return ws
}

func writeMessageToWebSocket(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over the websocket: %v", err)
	}

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

func assertStatus(t testing.TB, response *httptest.ResponseRecorder, want int) {
	t.Helper()
	got := response.Result().StatusCode
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
