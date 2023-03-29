package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.Reader
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	var league []Player
	_ = json.NewDecoder(f.database).Decode(&league)
	return league
}
