package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	app "github.com/nahcnuj/til/learn-go-with-tests/my-app"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := app.InMemoryPlayerStore{}
	server := app.NewServer(&store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func newGetScoreRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
}

func newPostWinRequest(player string) *http.Request {
	return httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
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
