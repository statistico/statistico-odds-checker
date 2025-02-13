package sportmonks

import (
	spClient "github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_parseExchangeMarketOdds(t *testing.T) {
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

func Test_convertOddsToRunners(t *testing.T) {
	t.Run("converts odds to runners for TEAM_CARDS market", func(t *testing.T) {
		t.Helper()

		totalOne := "Over 2.5"
		totalTwo := "Under 1.5"

		odds := []spClient.Odds{
			{
				Label: "1",
				Value: "3.75",
				Total: &totalOne,
			},
			{
				Label: "2",
				Value: "2.00",
				Total: &totalTwo,
			},
		}

		runners, err := convertOddsToRunners(odds, "TEAM_CARDS")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Len(t, runners, 2)

		assert.Equal(t, "HOME", runners[0].Name)
		assert.Equal(t, "OVER 2.5", *runners[0].Label)
		assert.Equal(t, float32(3.75), runners[0].BackPrices[0].Price)
		assert.Equal(t, "AWAY", runners[1].Name)
		assert.Equal(t, "UNDER 1.5", *runners[1].Label)
		assert.Equal(t, float32(2.00), runners[1].BackPrices[0].Price)
	})

	t.Run("converts odds to runners for BOTH_TEAMS_TO_SCORE market", func(t *testing.T) {
		t.Helper()

		labelOne := "1"
		labelTwo := "2"

		odds := []spClient.Odds{
			{
				Label: "Home",
				Value: "3.75",
				Name:  &labelOne,
			},
			{
				Label: "Away",
				Value: "3.75",
				Name:  &labelTwo,
			},
		}

		runners, err := convertOddsToRunners(odds, "BOTH_TEAMS_TO_SCORE")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Len(t, runners, 2)
	})

	t.Run("converts odds to runners for PLAYER_TO_SCORE_ANYTIME market", func(t *testing.T) {
		t.Helper()

		labelOne := "Mohammed Kudus"
		labelTwo := "Mo Salah"
		labelThree := "Cole Palmer"

		odds := []spClient.Odds{
			{
				Label: "Anytime",
				Value: "19",
				Name:  &labelOne,
			},
			{
				Label: "Anytime",
				Value: "3.75",
				Name:  &labelTwo,
			},
			{
				Label: "First",
				Value: "13.75",
				Name:  &labelThree,
			},
		}

		runners, err := convertOddsToRunners(odds, "PLAYER_TO_SCORE_ANYTIME")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Len(t, runners, 2)

		assert.Equal(t, "Mohammed Kudus", runners[0].Name)
		assert.Equal(t, "ANYTIME", *runners[0].Label)
		assert.Equal(t, float32(19.00), runners[0].BackPrices[0].Price)
		assert.Equal(t, "Mo Salah", runners[1].Name)
		assert.Equal(t, "ANYTIME", *runners[1].Label)
		assert.Equal(t, float32(3.75), runners[1].BackPrices[0].Price)
	})

	t.Run("converts odds to runners for PLAYER_TOTAL_SHOTS market", func(t *testing.T) {
		t.Helper()

		nameOne := "Mohammed Kudus"
		nameTwo := "Mo Salah"
		nameThree := "Cole Palmer"

		odds := []spClient.Odds{
			{
				Label:             "0.5",
				Value:             "2.75",
				Name:              &nameOne,
				MarketDescription: "Player Shots",
			},
			{
				Label:             "1.5",
				Value:             "3.75",
				Name:              &nameTwo,
				MarketDescription: "Player Shots",
			},
			{
				Label:             "1.5",
				Value:             "13.75",
				Name:              &nameThree,
				MarketDescription: "Player Shots On Target",
			},
		}

		runners, err := convertOddsToRunners(odds, "PLAYER_SHOTS_TOTAL")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Len(t, runners, 2)

		assert.Equal(t, "Mohammed Kudus", runners[0].Name)
		assert.Equal(t, "OVER 0.5", *runners[0].Label)
		assert.Equal(t, float32(2.75), runners[0].BackPrices[0].Price)
		assert.Equal(t, "Mo Salah", runners[1].Name)
		assert.Equal(t, "OVER 1.5", *runners[1].Label)
		assert.Equal(t, float32(3.75), runners[1].BackPrices[0].Price)
	})

	t.Run("converts odds to runners for MATCH_GOALS market", func(t *testing.T) {
		t.Helper()

		totalOne := "0.5"
		totalTwo := "1.5"

		odds := []spClient.Odds{
			{
				Label: "Over",
				Value: "2.75",
				Total: &totalOne,
			},
			{
				Label: "Under",
				Value: "3.75",
				Total: &totalTwo,
			},
			{
				Label: "Over",
				Value: "13.75",
				Total: &totalTwo,
			},
		}

		runners, err := convertOddsToRunners(odds, "MATCH_GOALS")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Len(t, runners, 3)

		assert.Equal(t, "OVER 0.5", runners[0].Name)
		assert.Equal(t, float32(2.75), runners[0].BackPrices[0].Price)
		assert.Equal(t, "UNDER 1.5", runners[1].Name)
		assert.Equal(t, float32(3.75), runners[1].BackPrices[0].Price)
		assert.Equal(t, "OVER 1.5", runners[2].Name)
		assert.Equal(t, float32(13.75), runners[2].BackPrices[0].Price)
	})

	t.Run("converts odds to runners for PLAYER_CARDS booked market", func(t *testing.T) {
		t.Helper()

		nameOne := "Mohammed Kudus"
		nameTwo := "Mo Salah"
		nameThree := "Cole Palmer"

		odds := []spClient.Odds{
			{
				Label: "Booked",
				Value: "2.75",
				Name:  &nameOne,
			},
			{
				Label: "1st Card",
				Value: "3.75",
				Name:  &nameTwo,
			},
			{
				Label: "Booked",
				Value: "13.75",
				Name:  &nameThree,
			},
		}

		runners, err := convertOddsToRunners(odds, "PLAYER_CARDS")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Len(t, runners, 2)

		assert.Equal(t, "Mohammed Kudus", runners[0].Name)
		assert.Equal(t, "BOOKED", *runners[0].Label)
		assert.Equal(t, float32(2.75), runners[0].BackPrices[0].Price)
		assert.Equal(t, "Cole Palmer", runners[1].Name)
		assert.Equal(t, "BOOKED", *runners[1].Label)
		assert.Equal(t, float32(13.75), runners[1].BackPrices[0].Price)
	})
}
