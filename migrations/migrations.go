package migrations

import (
	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/models"
)

func Migrate() {
	database.DBCon.AutoMigrate(models.Players{}, models.Matches{}, models.MatchData{}, models.Goals{})
}
