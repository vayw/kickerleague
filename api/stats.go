package api

import (
	"database/sql"
	"fmt"
	"sort"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vayw/kickerleague/database"
	"github.com/vayw/kickerleague/game"
	"github.com/vayw/kickerleague/models"
)

type Goal struct {
	ID   int `gorm:"column:player_id"`
	Ts   time.Time
	Auto bool
}

type MResult struct {
	Red    int
	Blue   int
	Lineup []game.Player
	Goals  []Goal
}

type MR struct {
	Result []MResult `json:"result"`
}

// @Summary Get match results
// @Description get match results
// @ID match-results
// @Accept  json
// @Produce  json
// @Param   num body int true "Number of match results"
// @Param   from body string false "Start date"
// @Param   to body string false "End date"
// @Success 200 {object} MR
// @Router /api/stats/matchresults [post]
func matchResults(c *gin.Context) {
	type Data struct {
		Num  int    `json:"num"`
		From string `json:"from"`
		To   string `json:"to"`
	}

	var data Data
	var matches []models.Match

	c.BindJSON(&data)

	result := make([]MResult, data.Num)
	database.DBCon.Limit(data.Num).Order("ts desc").Find(&matches)

	var matchlineup []models.MatchData
	for index, i := range matches {
		result[index] = MResult{}
		result[index].Red = i.Red_score
		result[index].Blue = i.Blue_score
		result[index].Lineup = make([]game.Player, 4)
		database.DBCon.Where("match_id = ?", i.ID).Find(&matchlineup)
		for pindex, j := range matchlineup {
			result[index].Lineup[pindex] = game.Player{j.PlayerID, j.Team, j.Position}
		}
		database.DBCon.Table("goals").Select("player_id, auto, ts").Where("match_id = ?", i.ID).Scan(&result[index].Goals)
	}
	r := MR{result}
	c.JSON(200, r)
}

type Scorer struct {
	Total int
	Id    int `gorm:"column:player_id"`
	Games int
}

type Scorers struct {
	Result   []Scorer `json:"result"`
	Position string   `json:"position"`
	Auto     bool     `json:"auto"`
}

// @Summary Get scorers table
// @Description goals scored by each player with games count
// @ID scorers-table
// @Accept  json
// @Produce  json
// @Param   position body string false "position to search for"
// @Param   from body string false "Start date"
// @Param   to body string false "End date"
// @Param   auto body string false "autogoals"
// @Success 200 {object} Scorers
// @Router /api/stats/ratings/goals [post]
func scorersTable(c *gin.Context) {
	type Data struct {
		Position string `json:"position"`
		From     string `json:"from"`
		To       string `json:"to"`
		Auto     string `json:"auto"`
	}
	var data Data
	c.BindJSON(&data)

	var results Scorers
	var rows *sql.Rows

	var isauto bool
	switch data.Auto {
	case "True", "true":
		isauto = true
	default:
		isauto = false
	}
	results.Position = data.Position
	results.Auto = isauto

	switch data.Position {
	case "Defender", "Forward":
		database.DBCon.Table("goals").Select("goals.player_id, count(*) as total").Joins("join match_data on match_data.match_id=goals.match_id and match_data.player_id = goals.player_id").Where("auto=? and position=?", isauto, data.Position).Group("goals.player_id").Order("total desc").Scan(&results.Result)
		rows, _ = database.DBCon.Table("match_data").Select("count(*), player_id").Where("position=?", data.Position).Group("player_id").Rows()
	default:
		data.Position = "overall"
		database.DBCon.Table("goals").Select("player_id, count(*) as total").Where("auto=?", isauto).Group("player_id").Order("total desc").Scan(&results.Result)
		rows, _ = database.DBCon.Table("match_data").Select("count(*), player_id").Group("player_id").Rows()
	}
	var count int
	var player_id int
	for rows.Next() {
		if err := rows.Scan(&count, &player_id); err != nil {
			fmt.Println(err)
		}
		for ind, scorer := range results.Result {
			if scorer.Id == player_id {
				results.Result[ind].Games = count
			}
		}
	}
	c.JSON(200, results)
}

type OverallResult struct {
	Matches int
	Goals   int
}

// @Summary Get overall stats
// @Description overall count of matches and goals
// @ID overall-stats
// @Accept  json
// @Produce  json
// @Success 200 {object} OverallResult
// @Router /api/stats/overall [get]
func overallStats(c *gin.Context) {
	var result OverallResult
	database.DBCon.Table("matches").Select("count(*) as matches").Scan(&result)
	database.DBCon.Table("goals").Select("count(*) as goals").Scan(&result)
	c.JSON(200, result)
}

type Stat struct {
	Id      int
	WinRate float32
	Wins    int
	Losses  int
}

type wrData struct {
	Player_id int
	Winner    string
	Team      string
	Position  string
}

type wrResults struct {
	Overall   []Stat
	Forwards  []Stat
	Defenders []Stat
}

// @Summary Get win rate stats
// @Description overall, defenders, forwards winrate
// @ID winrate-statss
// @Accept  json
// @Produce  json
// @Success 200 {object} wrResults
// @Router /api/stats/winrate [get]
func winRate(c *gin.Context) {

	var data []wrData
	win_map := make(map[int]int)
	defeat_map := make(map[int]int)
	forward_win_map := make(map[int]int)
	forward_defeat_map := make(map[int]int)
	defender_win_map := make(map[int]int)
	defender_defeat_map := make(map[int]int)

	database.DBCon.Table("match_data").Select("match_data.player_id, match_data.team, match_data.position, matches.winner").Joins("inner join matches on match_data.match_id = matches.id").Scan(&data)
	for _, row := range data {
		if row.Winner == row.Team {
			win_map[row.Player_id] += 1
			switch row.Position {
			case "Forward":
				forward_win_map[row.Player_id] += 1
			case "Defender":
				defender_win_map[row.Player_id] += 1
			}
		} else {
			defeat_map[row.Player_id] += 1
			switch row.Position {
			case "Forward":
				forward_defeat_map[row.Player_id] += 1
			case "Defender":
				defender_defeat_map[row.Player_id] += 1
			}
		}
	}

	overall := CalcAndSort(win_map, defeat_map)
	forwards := CalcAndSort(forward_win_map, forward_defeat_map)
	defenders := CalcAndSort(defender_win_map, defender_defeat_map)

	result := wrResults{overall, forwards, defenders}
	c.JSON(200, result)
}

func CalcAndSort(wins map[int]int, defeats map[int]int) []Stat {
	var result []Stat
	for key, value := range wins {
		result = append(result, Stat{key, float32(value) / float32(defeats[key]+value),
			value, defeats[key]})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].WinRate > result[j].WinRate
	})
	return result
}

type pmData struct {
}

type pmResult struct {
	Result []pmData
}

// @Summary Get player matches
// @Description Return players list of matches
// @ID players-matches
// @Accept  json
// @Produce  json
// @Param   position body string false "position to search for"
// @Param   from body string false "Start date"
// @Param   to body string false "End date"
// @Success 200 {object} pmResult
// @Router /api/stats/player/matches [post]
