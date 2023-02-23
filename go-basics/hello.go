package main

import (
	"fmt"
	"math"
)

type Batter struct {
	PlayerID             string
	YearID               int
	Stint                int
	TeamID               string
	LeagueID             string
	Games                int
	AtBats               int
	Runs                 int
	Hits                 int
	Doubles              int
	Triples              int
	Homeruns             int
	RunsBattedIn         int
	StolenBases          int
	CaughtStealing       int
	Walks                int
	Strikeouts           int
	IntentionalWalks     int
	HitByPitch           int
	SacrificeBunt        int
	SacrificeFly         int
	GroundIntoDoublePlay int
	BattingAverage       float64
	OnBasePercentage     float64
	SluggingPercentage   float64
	OnBasePlusSlugging   float64
}

func toFixed(stat float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(int(stat*output)) / output
}

func getAverages(b *Batter) {
	getBattingAvg(b)
	getOnBasePercentage(b)
	getSluggingPercentage(b)
	getOnBasePlusSlugging(b)
}

func getBattingAvg(b *Batter) {
	b.BattingAverage = toFixed(float64(b.Hits)/float64(b.AtBats), 3)
}

func getOnBasePercentage(b *Batter) {
	b.OnBasePercentage = toFixed(float64(b.Hits+b.Walks+b.IntentionalWalks+b.HitByPitch)/float64(b.AtBats+b.Walks+b.IntentionalWalks+b.SacrificeFly), 3)
}

func getSluggingPercentage(b *Batter) {
	b.SluggingPercentage = toFixed(float64(b.Hits+b.Doubles+(2*b.Triples)+(3*b.Homeruns))/float64(b.AtBats), 3)
}

func getOnBasePlusSlugging(b *Batter) {
	b.OnBasePlusSlugging = toFixed(b.OnBasePercentage+b.SluggingPercentage, 3)
}

func main() {
	playerOne := Batter{
		PlayerID:             "troutmi01",
		YearID:               2014,
		Stint:                1,
		TeamID:               "LAA",
		LeagueID:             "AL",
		Games:                157,
		AtBats:               602,
		Runs:                 115,
		Hits:                 173,
		Doubles:              39,
		Triples:              9,
		Homeruns:             36,
		RunsBattedIn:         111,
		StolenBases:          16,
		CaughtStealing:       2,
		Walks:                83,
		Strikeouts:           184,
		IntentionalWalks:     6,
		HitByPitch:           10,
		SacrificeBunt:        0,
		SacrificeFly:         10,
		GroundIntoDoublePlay: 6,
	}

	playerTwo := Batter{
		PlayerID:             "cabremi01",
		YearID:               2013,
		Stint:                1,
		TeamID:               "DET",
		LeagueID:             "AL",
		Games:                148,
		AtBats:               555,
		Runs:                 103,
		Hits:                 193,
		Doubles:              26,
		Triples:              1,
		Homeruns:             44,
		RunsBattedIn:         137,
		StolenBases:          3,
		CaughtStealing:       0,
		Walks:                90,
		Strikeouts:           94,
		IntentionalWalks:     19,
		HitByPitch:           5,
		SacrificeBunt:        0,
		SacrificeFly:         2,
		GroundIntoDoublePlay: 19,
	}

	playerThree := Batter{
		PlayerID:             "hamiljo03",
		YearID:               2010,
		Stint:                1,
		TeamID:               "TEX",
		LeagueID:             "AL",
		Games:                133,
		AtBats:               518,
		Runs:                 95,
		Hits:                 186,
		Doubles:              40,
		Triples:              3,
		Homeruns:             32,
		RunsBattedIn:         100,
		StolenBases:          8,
		CaughtStealing:       1,
		Walks:                43,
		Strikeouts:           95,
		IntentionalWalks:     5,
		HitByPitch:           5,
		SacrificeBunt:        1,
		SacrificeFly:         4,
		GroundIntoDoublePlay: 11,
	}

	playerFour := Batter{
		PlayerID:             "mauerjo01",
		YearID:               2009,
		Stint:                1,
		TeamID:               "MIN",
		LeagueID:             "AL",
		Games:                138,
		AtBats:               523,
		Runs:                 94,
		Hits:                 191,
		Doubles:              30,
		Triples:              1,
		Homeruns:             28,
		RunsBattedIn:         96,
		StolenBases:          4,
		CaughtStealing:       1,
		Walks:                76,
		Strikeouts:           63,
		IntentionalWalks:     14,
		HitByPitch:           2,
		SacrificeBunt:        0,
		SacrificeFly:         5,
		GroundIntoDoublePlay: 13,
	}

	fmt.Println("Play Ball!")
	getAverages(&playerOne)
	getAverages(&playerTwo)
	getAverages(&playerThree)
	getAverages(&playerFour)
	fmt.Println(playerOne)
	fmt.Println(playerTwo)
	fmt.Println(playerThree)
	fmt.Println(playerFour)
}
