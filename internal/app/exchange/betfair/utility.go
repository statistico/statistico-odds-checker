package betfair

import (
	"github.com/statistico/statistico-betfair-go-client"
)

const (
	Away      = "Away"
	Draw      = "Draw"
	Home      = "Home"
	MatchOdds = "MATCH_ODDS"
)

func parseRunnerName(runner *betfair.RunnerCatalogue, market string) string {
	if market != MatchOdds {
		return runner.RunnerName
	}

	if runner.SortPriority == 1 {
		return Home
	}

	if runner.SortPriority == 2 {
		return Away
	}

	return Draw
}
