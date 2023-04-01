package betfair

import (
	"github.com/statistico/statistico-betfair-go-client"
)

const (
	Away      = "Away"
	Draw      = "Draw"
	Home      = "Home"
	MatchOdds = "MATCH_ODDS"
	Over      = "OVER"
	Under     = "UNDER"
)

func parseRunnerName(runner *betfair.RunnerCatalogue, market string) string {
	if market == MatchOdds {
		return handleMatchOddsRunners(runner.SortPriority)
	}

	if runner.RunnerName == "Under 2.5 Goals" {
		return Under
	}

	if runner.RunnerName == "Over 2.5 Goals" {
		return Over
	}

	return runner.RunnerName
}

func handleMatchOddsRunners(priority int) string {
	if priority == 1 {
		return Home
	}

	if priority == 2 {
		return Away
	}

	return Draw
}
