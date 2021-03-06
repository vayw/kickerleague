package game

import (
	"errors"
	"os"
	"testing"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/migrations"
	"github.com/vayw/kickerleague/player"
)

func initMatch() error {
	database.ConnectDB()
	migrations.Migrate(database.DBCon)
	pl := []string{"a", "b", "c", "d"}
	for _, n := range pl {
		_, err := player.AddPlayer(n)
		if err != nil {
			if err.Error() != "player already exists" {
				errors.New("error creating test players")
			}
		}
	}
	return nil
}

func TestNewMatch(t *testing.T) {
	if err := initMatch(); err != nil {
		t.Error(err)
	}
	lineup := []Player{
		Player{1, "Red", "Defender"},
		Player{2, "Red", "Forward"},
		Player{3, "Blue", "Defender"},
		Player{4, "Blue", "Forward"},
	}
	_, err := NewMatch(lineup)
	if err != nil {
		t.Error(err)
	}
	defer database.DBCon.Close()
	Clean()
}

func TestScore(t *testing.T) {
	if err := initMatch(); err != nil {
		t.Error(err)
	}
	lineup := []Player{
		Player{1, "Red", "Defender"},
		Player{2, "Red", "Forward"},
		Player{3, "Blue", "Defender"},
		Player{4, "Blue", "Forward"},
	}
	mid, err := NewMatch(lineup)
	if err != nil {
		t.Error(err)
	}
	if err := Score(1, mid, false); err != nil {
		t.Error(err)
	}
	defer database.DBCon.Close()
	Clean()
}

func TestEndMatch(t *testing.T) {
	if err := initMatch(); err != nil {
		t.Error(err)
	}
	lineup := []Player{
		Player{1, "Red", "Defender"},
		Player{2, "Red", "Forward"},
		Player{3, "Blue", "Defender"},
		Player{4, "Blue", "Forward"},
	}

	mid, err := NewMatch(lineup)
	if err != nil {
		t.Error(err)
	}
	for _, i := range [7]int{1, 3, 3, 4, 2, 3, 3} {
		if err := Score(i, mid, false); err != nil {
			t.Error(err)
		}
	}
	if err := Score(4, mid, true); err != nil {
		t.Error(err)
	}
	res, err := EndMatch(mid)
	if err != nil {
		t.Error(err)
	}
	defer database.DBCon.Close()
	if res.Red != 3 || res.Blue != 5 || res.Winner != "Blue" {
		t.Error(res)
	}
	Clean()
}

func Clean() {
	os.Remove("kicker.db")
}
