package player

import (
	"errors"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/models"
)

func AddPlayer(name string) (models.Player, error) {
	player := models.Player{Name: name}
	database.InitDB()
	defer database.DBCon.Close()

	if database.DBCon.NewRecord(player) {
		database.DBCon.Create(&player)
		return player, nil
	} else {
		return models.Player{}, errors.New("player already exists")
	}
}
