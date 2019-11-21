package models

import "time"

type Players struct {
	ID   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
}

type Matches struct {
	ID         int `gorm:"primary_key;AUTO_INCREMENT"`
	Red_score  int
	Blue_score int
	Winner     string
	TS         time.Time
}

type MatchData struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	Position string
	Team     string
	Matches  Matches `gorm:"foreignkey:MatchID"`
	MatchID  int
}

type Goals struct {
	ID       int     `gorm:"primary_key;AUTO_INCREMENT"`
	Players  Players `gorm:"foreignkey:PlayerID"`
	PlayerID int
	Matches  Matches `gorm:"foreignkey:MatchID"`
	MatchID  int
}
