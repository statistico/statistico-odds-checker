package exchange_test

import (
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertOddsToRunners(t *testing.T) {
	t.Run("converts sportsmonks.Odds slice to Runner slice", func(t *testing.T) {
		t.Helper()

		odds := []sportmonks.Odds{
			{
				Label: "Over",
				Dp3:   "1.98",
			},
			{
				Label: "Under",
				Dp3:   "2.55",
			},
		}

		runners, err := exchange.ConvertOddsToRunners(odds)

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		expected := []*exchange.Runner{
			{
				ID:   0,
				Name: "OVER",
				BackPrices: []exchange.PriceSize{
					{
						Price: 1.98,
						Size:  0,
					},
				},
			},
			{
				ID:   0,
				Name: "UNDER",
				BackPrices: []exchange.PriceSize{
					{
						Price: 2.55,
						Size:  0,
					},
				},
			},
		}

		assert.Equal(t, expected, runners)
	})

	t.Run("return an error if price value is not a valid float", func(t *testing.T) {
		t.Helper()

		odds := []sportmonks.Odds{
			{
				Label: "Over",
				Dp3:   "1.98",
			},
			{
				Label: "Under",
				Dp3:   "HeeeHeee",
			},
		}

		_, err := exchange.ConvertOddsToRunners(odds)

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "value 'HeeeHeee' is not a valid floating point number", err.Error())
	})
}
