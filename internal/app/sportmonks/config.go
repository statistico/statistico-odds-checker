package sportmonks

var markets = map[string][]string{
	"3Way Result": {
		"MATCH_ODDS",
	},
	"Over/Under": {
		"OVER_UNDER_05",
		"OVER_UNDER_15",
		"OVER_UNDER_25",
		"OVER_UNDER_35",
		"OVER_UNDER_45",
	},
	"Both Teams To Score": {
		"BOTH_TEAMS_TO_SCORE",
	},
	"Corners Over Under": {
		"OVER_UNDER_55_CORNR",
		"OVER_UNDER_85_CORNR",
		"OVER_UNDER_95_CORNR",
		"OVER_UNDER_105_CORNR",
		"OVER_UNDER_115_CORNR",
		"OVER_UNDER_125_CORNR",
		"OVER_UNDER_135_CORNR",
	},
}
