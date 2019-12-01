package player

import (
	"os"
	"testing"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/migrations"
)

func TestAddPlayer(t *testing.T) {
	database.ConnectDB()
	migrations.Migrate()
	defer database.DBCon.Close()

	_, err := AddPlayer("Obi-Van")
	os.Remove("kicker.db")
	if err != nil {
		if err.Error() != "player already exists" {
			t.Error(err)
		}
	}
}
