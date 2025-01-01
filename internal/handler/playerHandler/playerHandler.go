package playerhandler

import (
	"fmt"
	"net/http"
)

func NewPlayerHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/players/{id}", getPlayer)
	return mux
}

func getPlayer(w http.ResponseWriter, r *http.Request) {
	player := r.PathValue("id")

	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprint(w, "10")
		return
	}
}
