package pinnacle

import (
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-sportmonks-go-client"
	"strconv"
)

func ConvertOddsToRunners(odds []sportmonks.Odds) ([]*exchange.Runner, error) {
	runners := []*exchange.Runner{}

	for _, o := range odds {
		price, err := strconv.ParseFloat(o.Dp3, 32)

		if err != nil {
			return nil, fmt.Errorf("value '%s' is not a valid decimal", o.Dp3)
		}

		runners = append(runners, &exchange.Runner{
			Name: o.Label,
			BackPrices: []exchange.PriceSize{
				{
					Price: float32(price),
					Size:  0,
				},
			},
		})
	}

	return runners, nil
}
