package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBCon *gorm.DB
)

func ConnectDB() {
	var err error

	DBCon, err = gorm.Open("sqlite3", "kicker.db")

	if err != nil {
		panic("failed to connect database")
	}
}
