package models

import "time"

type Player struct {
	ID   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
}

type Match struct {
	ID         int `gorm:"primary_key;AUTO_INCREMENT"`
	Red_score  int
	Blue_score int
	Winner     string
	TS         time.Time
}

type MatchData struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Player   Player `gorm:"foreignkey:PlayerID"`
	PlayerID int
	Position string
	Team     string
	Match    Match `gorm:"foreignkey:MatchID"`
	MatchID  int
}

type Goal struct {
	ID       int    `gorm:"primary_key;AUTO_INCREMENT"`
	Player   Player `gorm:"foreignkey:PlayerID"`
	PlayerID int
	Match    Match `gorm:"foreignkey:MatchID"`
	MatchID  int
}
