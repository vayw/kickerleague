package database

import "github.com/jinzhu/gorm"

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
