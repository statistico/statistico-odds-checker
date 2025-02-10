package sportmonks

import (
	spClient "github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_parseExchangeMarketOdds(t *testing.T) {
	//t.Run("parses odds for both teams to score market", func(t *testing.T) {
	//	t.Helper()
	//
	//	odds := []sportmonks.Odds{
	//		{
	//			Value:            &f,
	//			Handicap:         nil,
	//			Total:            "",
	//			Label:            "Yes",
	//			Probability:      "35.71%",
	//			Dp3:              "2.800",
	//			American:         0,
	//			Fractional:       nil,
	//			Winning:          nil,
	//			Stop:             false,
	//			BookmakerEventID: nil,
	//			LastUpdate:       sportmonks.DateTime{},
	//		},
	//	}
	//
	//	parsed, err := parseMarketRunners("BOTH_TEAMS_TO_SCORE", 70, odds)
	//
	//	if err != nil {
	//		t.Fatalf("Expected nil, got %s", err.Error())
	//	}
	//
	//	assert.Equal(t, odds, parsed)
	//})

	t.Run("parses odds for match odds market", func(t *testing.T) {
		t.Helper()

		labelOne := "1"
		labelTwo := "2"
		created, _ := time.Parse(time.RFC3339, "2025-01-26T14:57:19.000000Z")

		odds := []spClient.Odds{
			{
				ID:                    151577019200,
				FixtureID:             19155301,
				MarketID:              1,
				BookmakerID:           2,
				Label:                 "Home",
				Value:                 "3.75",
				Name:                  &labelOne,
				SortOrder:             nil,
				MarketDescription:     "Full Time Result",
				Probability:           "26.67%",
				DP3:                   "1.750",
				Fractional:            "15/4",
				American:              "275",
				Winning:               false,
				Stopped:               false,
				Total:                 nil,
				Handicap:              nil,
				Participants:          nil,
				CreatedAt:             created,
				OriginalLabel:         nil,
				LatestBookmakerUpdate: "2025-02-10 14:10:51",
				Bookmaker: &spClient.Bookmaker{
					ID:       2,
					LegacyID: 2,
					Name:     "bet365",
				},
				Market: &spClient.Market{
					ID:                     1,
					LegacyID:               1,
					Name:                   "Fulltime Result",
					DeveloperName:          "FULLTIME_RESULT",
					HasWinningCalculations: false,
				},
			},
			{
				ID:                    151577019200,
				FixtureID:             19155301,
				MarketID:              1,
				BookmakerID:           2,
				Label:                 "Away",
				Value:                 "3.75",
				Name:                  &labelTwo,
				SortOrder:             nil,
				MarketDescription:     "Full Time Result",
				Probability:           "26.67%",
				DP3:                   "3.750",
				Fractional:            "15/4",
				American:              "275",
				Winning:               false,
				Stopped:               false,
				Total:                 nil,
				Handicap:              nil,
				Participants:          nil,
				CreatedAt:             created,
				OriginalLabel:         nil,
				LatestBookmakerUpdate: "2025-02-10 14:10:51",
				Bookmaker: &spClient.Bookmaker{
					ID:       2,
					LegacyID: 2,
					Name:     "bet365",
				},
				Market: &spClient.Market{
					ID:                     1,
					LegacyID:               1,
					Name:                   "Fulltime Result",
					DeveloperName:          "FULLTIME_RESULT",
					HasWinningCalculations: false,
				},
			},
			{
				ID:                    151577019200,
				FixtureID:             19155301,
				MarketID:              1,
				BookmakerID:           5,
				Label:                 "Away",
				Value:                 "3.75",
				Name:                  &labelTwo,
				SortOrder:             nil,
				MarketDescription:     "Full Time Result",
				Probability:           "26.67%",
				DP3:                   "3.750",
				Fractional:            "15/4",
				American:              "275",
				Winning:               false,
				Stopped:               false,
				Total:                 nil,
				Handicap:              nil,
				Participants:          nil,
				CreatedAt:             created,
				OriginalLabel:         nil,
				LatestBookmakerUpdate: "2025-02-10 14:10:51",
				Bookmaker: &spClient.Bookmaker{
					ID:       5,
					LegacyID: 7,
					Name:     "888Bet",
				},
				Market: &spClient.Market{
					ID:                     1,
					LegacyID:               1,
					Name:                   "Fulltime Result",
					DeveloperName:          "FULLTIME_RESULT",
					HasWinningCalculations: false,
				},
			},
		}

		parsed := parseExchangeMarketOdds(2, odds)

		assert.Equal(t, []spClient.Odds{odds[0], odds[1]}, parsed)
	})
}
