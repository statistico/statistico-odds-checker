package sportmonks

import (
	"github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseMarketName(t *testing.T) {
	t.Run("successfully parses name for the given market", func(t *testing.T) {
		t.Helper()

		markets := []struct {
			MarketName string
			Market     string
		}{
			{
				MarketName: "MATCH_ODDS",
				Market:     "3Way Result",
			},
			{
				MarketName: "OVER_UNDER_05",
				Market:     "Over\\/Under",
			},
			{
				MarketName: "OVER_UNDER_15",
				Market:     "Over\\/Under",
			},
			{
				MarketName: "OVER_UNDER_25",
				Market:     "Over\\/Under",
			},
			{
				MarketName: "OVER_UNDER_35",
				Market:     "Over\\/Under",
			},
			{
				MarketName: "OVER_UNDER_45",
				Market:     "Over\\/Under",
			},
			{
				MarketName: "BOTH_TEAMS_TO_SCORE",
				Market:     "Both Teams To Score",
			},
			{
				MarketName: "OVER_UNDER_55_CORNR",
				Market:     "Corners Over Under",
			},
			{
				MarketName: "OVER_UNDER_85_CORNR",
				Market:     "Corners Over Under",
			},
			{
				MarketName: "OVER_UNDER_95_CORNR",
				Market:     "Corners Over Under",
			},
			{
				MarketName: "OVER_UNDER_105_CORNR",
				Market:     "Corners Over Under",
			},
			{
				MarketName: "OVER_UNDER_115_CORNR",
				Market:     "Corners Over Under",
			},
			{
				MarketName: "OVER_UNDER_125_CORNR",
				Market:     "Corners Over Under",
			},
			{
				MarketName: "OVER_UNDER_135_CORNR",
				Market:     "Corners Over Under",
			},
		}

		for _, m := range markets {
			marketId, err := parseMarketName(m.MarketName)

			if err != nil {
				t.Fatalf("Expected nil, got %s", err.Error())
			}

			assert.Equal(t, m.Market, marketId)
		}
	})

	t.Run("returns an error if market is not supported", func(t *testing.T) {
		t.Helper()

		_, err := parseMarketName("ASIAN_HANDICAP")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "market 'ASIAN_HANDICAP' is not supported", err.Error())
	})
}

func Test_parseMarketRunners(t *testing.T) {
	f := sportmonks.FlexFloat(2.80)

	t.Run("parses odds for both teams to score market", func(t *testing.T) {
		t.Helper()

		odds := []sportmonks.Odds{
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "",
				Label:            "Yes",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
		}

		parsed, err := parseMarketRunners("BOTH_TEAMS_TO_SCORE", 70, odds)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, odds, parsed)
	})

	t.Run("parses odds for match odds market", func(t *testing.T) {
		t.Helper()

		odds := []sportmonks.Odds{
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "",
				Label:            "X",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
		}

		parsed, err := parseMarketRunners("MATCH_ODDS", 70, odds)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, odds, parsed)
	})

	t.Run("parses odds for over under 25 market", func(t *testing.T) {
		t.Helper()

		odds := []sportmonks.Odds{
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "2.25",
				Label:            "Under",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "2.25",
				Label:            "Over",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "1.50",
				Label:            "Under",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
		}

		expected := []sportmonks.Odds{odds[0], odds[1]}

		parsed, err := parseMarketRunners("OVER_UNDER_25", 70, odds)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, expected, parsed)
	})

	t.Run("parses odds for over under 25 market using the alternative total", func(t *testing.T) {
		t.Helper()

		odds := []sportmonks.Odds{
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "2.5",
				Label:            "Under",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "2.5",
				Label:            "Over",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
			{
				Value:            &f,
				Handicap:         nil,
				Total:            "1.50",
				Label:            "Under",
				Probability:      "35.71%",
				Dp3:              "2.800",
				American:         0,
				Fractional:       nil,
				Winning:          nil,
				Stop:             false,
				BookmakerEventID: nil,
				LastUpdate:       sportmonks.DateTime{},
			},
		}

		expected := []sportmonks.Odds{odds[0], odds[1]}

		parsed, err := parseMarketRunners("OVER_UNDER_25", 70, odds)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		assert.Equal(t, expected, parsed)
	})
}
