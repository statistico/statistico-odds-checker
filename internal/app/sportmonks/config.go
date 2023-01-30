package pinnacle

import "fmt"

var markets = map[int][]string{
	1: {
		"MATCH_ODDS",
	},
	12: {
		"OVER_UNDER_05",
		"OVER_UNDER_15",
		"OVER_UNDER_25",
		"OVER_UNDER_35",
		"OVER_UNDER_45",
	},
	976105: {
		"BOTH_TEAMS_TO_SCORE",
	},
	976384: {
		"OVER_UNDER_55_CORNR",
		"OVER_UNDER_85_CORNR",
		"OVER_UNDER_95_CORNR",
		"OVER_UNDER_105_CORNR",
		"OVER_UNDER_115_CORNR",
		"OVER_UNDER_125_CORNR",
		"OVER_UNDER_135_CORNR",
	},
	136704529: {
		"MATCH_SHOTS_TARGET",
	},
	136704537: {
		"MATCH_SHOTS",
	},
}

func parseMarketId(market string) (int, error) {
	for id, m := range markets {
		for _, mk := range m {
			if mk == market {
				return id, nil
			}
		}
	}

	return 0, fmt.Errorf("market '%s' is not supported", market)
}
