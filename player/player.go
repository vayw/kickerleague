package player

import (
	"errors"

	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/models"
)

func AddPlayer(name string) (models.Player, error) {
	player := models.Player{Name: name}

	if err := database.DBCon.Create(&player).Error; err == nil {
		return player, nil
	} else {
		switch err.Error() {
		case "UNIQUE constraint failed: players.name":
			return models.Player{}, errors.New("player already exists")
		default:
			return models.Player{}, err
		}
	}
}

func PlayerList() ([]models.Player, error) {
	var Players []models.Player
	database.DBCon.Find(&Players)

	return Players, nil
}
