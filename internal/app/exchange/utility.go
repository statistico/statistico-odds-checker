package exchange

import (
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
	"strconv"
)

func ConvertOddsToRunners(odds []sportmonks.Odds) ([]*Runner, error) {
	runners := []*Runner{}

	for _, o := range odds {
		price, err := strconv.ParseFloat(o.Dp3, 32)

		if err != nil {
			return nil, fmt.Errorf("value '%s' is not a valid decimal", o.Dp3)
		}

		runners = append(runners, &Runner{
			Name: o.Label,
			BackPrices: []PriceSize{
				{
					Price: float32(price),
					Size:  0,
				},
			},
		})
	}

	return runners, nil
}
