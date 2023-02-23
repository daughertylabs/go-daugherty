package batting

import "math"

func toFixed(stat float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(int(stat*output)) / output
}

// Batting Average (AVG): H / AB (calculated to 3 decimal places)
func CalculateBattingAverage(batterPtr *Batter) {
	batterPtr.BattingAverage = toFixed(float64(batterPtr.Hits)/float64(batterPtr.AtBats), 3)
}
