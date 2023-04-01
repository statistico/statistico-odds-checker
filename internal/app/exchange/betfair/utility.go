package betfair

import (
	"github.com/statistico/statistico-betfair-go-client"
	"strings"
)

const (
	Away      = "Away"
	Draw      = "Draw"
	Home      = "Home"
	MatchOdds = "MATCH_ODDS"
	No        = "NO"
	Over      = "OVER"
	Under     = "UNDER"
	Yes       = "YES"
)

func parseRunnerName(runner *betfair.RunnerCatalogue, market string) string {
	if market == MatchOdds {
		return handleMatchOddsRunners(runner.SortPriority)
	}

	if strings.HasPrefix(runner.RunnerName, "Under") {
		return Under
	}

	if strings.HasPrefix(runner.RunnerName, "Over") {
		return Over
	}

	if runner.RunnerName == "No" {
		return No
	}

	if runner.RunnerName == "Yes" {
		return Yes
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
