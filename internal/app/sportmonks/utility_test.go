package sportmonks

import (
	"github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_parseMarketId(t *testing.T) {
	t.Run("successfully parses ID for the given market", func(t *testing.T) {
		t.Helper()

		markets := []struct {
			MarketName string
			MarketId   int
		}{
			{
				MarketName: "MATCH_ODDS",
				MarketId:   1,
			},
			{
				MarketName: "OVER_UNDER_05",
				MarketId:   12,
			},
			{
				MarketName: "OVER_UNDER_15",
				MarketId:   12,
			},
			{
				MarketName: "OVER_UNDER_25",
				MarketId:   12,
			},
			{
				MarketName: "OVER_UNDER_35",
				MarketId:   12,
			},
			{
				MarketName: "OVER_UNDER_45",
				MarketId:   12,
			},
			{
				MarketName: "BOTH_TEAMS_TO_SCORE",
				MarketId:   976105,
			},
			{
				MarketName: "OVER_UNDER_55_CORNR",
				MarketId:   976384,
			},
			{
				MarketName: "OVER_UNDER_85_CORNR",
				MarketId:   976384,
			},
			{
				MarketName: "OVER_UNDER_95_CORNR",
				MarketId:   976384,
			},
			{
				MarketName: "OVER_UNDER_105_CORNR",
				MarketId:   976384,
			},
			{
				MarketName: "OVER_UNDER_115_CORNR",
				MarketId:   976384,
			},
			{
				MarketName: "OVER_UNDER_125_CORNR",
				MarketId:   976384,
			},
			{
				MarketName: "OVER_UNDER_135_CORNR",
				MarketId:   976384,
			},
		}

		for _, m := range markets {
			marketId, err := parseMarketId(m.MarketName)

			if err != nil {
				t.Fatalf("Expected nil, got %s", err.Error())
			}

			assert.Equal(t, m.MarketId, marketId)
		}
	})

	t.Run("returns an error if market is not supported", func(t *testing.T) {
		t.Helper()

		_, err := parseMarketId("ASIAN_HANDICAP")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "market 'ASIAN_HANDICAP' is not supported", err.Error())
	})
}

func Test_parseExchangeOdds(t *testing.T) {
	f := sportmonks.FlexFloat(2.80)

	exchange := sportmonks.MatchOdds{
		ID:        1,
		Name:      "3Way Result",
		Suspended: false,
		BookmakerOddsData: sportmonks.BookmakerOddsData{Data: []sportmonks.BookmakerOdds{
			{
				ID:   70,
				Name: "Pncl",
				OddsData: sportmonks.OddsData{
					Data: []sportmonks.Odds{
						{
							Value:            &f,
							Handicap:         nil,
							Total:            "",
							Label:            "1",
							Probability:      "35.71%",
							Dp3:              "2.809",
							American:         179,
							Fractional:       nil,
							Winning:          nil,
							Stop:             false,
							BookmakerEventID: nil,
							LastUpdate:       sportmonks.DateTime{},
						},
					},
				},
			},
		}},
	}

	t.Run("parses odds for a given exchange ID", func(t *testing.T) {
		t.Helper()

		odds := parseExchangeOdds(70, []sportmonks.MatchOdds{exchange})

		assert.Equal(t, exchange.BookmakerOddsData.Data[0].Odds(), odds)
	})

	t.Run("returns nil if unable to parse exchange", func(t *testing.T) {
		t.Helper()

		odds := parseExchangeOdds(999, []sportmonks.MatchOdds{exchange})

		if odds != nil {
			t.Fatalf("Expected nil, got %+v", odds)
		}
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
