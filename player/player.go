package player

import (
	"errors"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/models"
)

func AddPlayer(name string) (models.Player, error) {
	player := models.Player{Name: name}
	database.ConnectDB()
	defer database.DBCon.Close()

	if err := database.DBCon.Create(&player).Error; err == nil {
		return player, nil
	} else {
		return models.Player{}, errors.New("player already exists")
	}
}
