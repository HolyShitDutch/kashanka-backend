package main

import (
	"log"
	"net/http"

	"github.com/moroz/kashanka/handlers"
)

func main() {
	r := handlers.Router()

	log.Print("Listening on :3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
