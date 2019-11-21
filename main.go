package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/cli"
	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/migrations"
)

func main() {
	app := &cli.App{
		Name:  "help",
		Usage: "about",
		Action: func(c *cli.Context) error {
			fmt.Println("this is a dummy kicker league server")
			return nil
		},
	}

	database.InitDB()
	migrations.Migrate()
	defer database.DBCon.Close()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
