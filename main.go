package main

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/vayw/kickerleague/cmd"
)

func main() {
	cmd.Execute()
}
