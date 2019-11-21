package database

import "github.com/jinzhu/gorm"

var (
	DBCon *gorm.DB
)

func InitDB() {
	var err error

	DBCon, err = gorm.Open("sqlite3", "kicker.db")

	if err != nil {
		panic("failed to connect database")
	}

}
