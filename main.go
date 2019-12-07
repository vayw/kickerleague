package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vayw/kickerleague/cmd"
	"github.com/vayw/kickerleague/database"
)

func main() {
	database.ConnectDB()
	defer database.DBCon.Close()
	cmd.Execute()
}
