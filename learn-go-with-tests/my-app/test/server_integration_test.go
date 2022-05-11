package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	app "github.com/nahcnuj/til/learn-go-with-tests/my-app"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "")
	defer cleanDatabase()
	store := app.NewFileSystemPlayerStore(database)
	server := app.NewServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response.Result(), "application/json")

		got := getLeagueFromResponse(t, response.Body)
		want := []app.Player{
			{Name: "Pepper", Wins: 3},
		}
		assertLeague(t, got, want)
	})
}

func newGetScoreRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
}

func newPostWinRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
}

func newGetLeagueRequest() *http.Request {
	return httptest.NewRequest(http.MethodGet, "/league", nil)
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league []app.Player) {
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

func assertLeague(t testing.TB, got, want []app.Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("league table is wrong, got %v, want %v", got, want)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Errorf("could not create temp file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
