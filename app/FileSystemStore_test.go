package poker_test

import (
	"os"
	"testing"

	"github.com/igorakimy/poker"
)

func TestFileSystemStore(t *testing.T) {

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)

		poker.AssertNoError(t, err)

		got := store.GetLeague()

		want := []poker.Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("/get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		got := store.GetPlayerScore("Chris")
		want := 33
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}
		]`)
		defer cleanDatabase()

		store, err := poker.NewFileSystemPlayerStore(database)
		poker.AssertNoError(t, err)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1
		poker.AssertScoreEquals(t, got, want)
	})

	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := poker.NewFileSystemPlayerStore(database)

		poker.AssertNoError(t, err)
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file, %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}
