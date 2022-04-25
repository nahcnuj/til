package my_app

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, getPlayerScore(player))
}

func getPlayerScore(player string) string {
	if player == "Pepper" {
		return "20"
	}
	if player == "Floyd" {
		return "10"
	}
	return ""
}
