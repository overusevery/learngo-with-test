package main

import (
	"log"
	playerhandler "mypractice/httpserver/internal/handler/playerHandler"
	"net/http"
)

func main() {
	mux := playerhandler.NewPlayerHandler()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
