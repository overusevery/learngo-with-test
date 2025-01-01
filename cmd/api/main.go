package main

import (
	"log"
	playerhandler "mypractice/httpserver/internal/handler/playerHandler"
	"net/http"
)

func main() {
	// Hello world, the web server

	http.HandleFunc("/api/players", playerhandler.PlayerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
