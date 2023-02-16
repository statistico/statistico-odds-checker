package exchange

import (
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
	"strconv"
	"strings"
)

func ConvertOddsToRunners(odds []sportmonks.Odds) ([]*Runner, error) {
	var runners []*Runner

	for _, o := range odds {
		price, err := strconv.ParseFloat(o.Dp3, 32)

		if err != nil {
			return nil, fmt.Errorf("value '%s' is not a valid floating point number", o.Dp3)
		}

		runners = append(runners, &Runner{
			Name: strings.ToUpper(o.Label),
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
