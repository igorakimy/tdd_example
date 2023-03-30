package main

import (
	"log"
	"net/http"

	"github.com/igorakimy/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFile, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer closeFile()

	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000: %v", err)
	}
}
