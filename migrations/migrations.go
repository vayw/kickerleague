package migrations

import (
	"time"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func Migrate(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "initial",
			Migrate: func(tx *gorm.DB) error {
				type Player struct {
					ID   int    `gorm:"primary_key;AUTO_INCREMENT"`
					Name string `gorm:"unique;not null"`
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
					TS       time.Time
				}

				return tx.AutoMigrate(Player{}, Match{}, MatchData{}, Goal{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("player", "matches", "match_data", "goals").Error
			},
		},
		{
			ID: "autogoals",
			Migrate: func(tx *gorm.DB) error {
				type Goal struct {
					Auto bool `gorm:"DEFAULT:false"`
				}
				return tx.AutoMigrate(Goal{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Table("goals").DropColumn("Auto").Error
			},
		},
	})
	return m.Migrate()
}
