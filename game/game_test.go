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
	migrations.Migrate()
	defer database.DBCon.Close()
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
	os.Remove("kicker.db")
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
	if err := Score(1, mid); err != nil {
		t.Error(err)
	}
	os.Remove("kicker.db")
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
		if err := Score(i, mid); err != nil {
			t.Error(err)
		}
	}
	res, err := EndMatch(mid)
	if err != nil {
		t.Error(err)
	}
	os.Remove("kicker.db")
	if res.Red != 2 || res.Blue != 5 || res.Winner != "blue" {
		t.Error(res)
	}
}
