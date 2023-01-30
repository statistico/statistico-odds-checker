package pinnacle

import (
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-sportmonks-go-client"
	"strconv"
)

func ParseMarketRunners(market string, exchangeID uint64, odds []sportmonks.Odds) ([]sportmonks.Odds, error) {
	switch market {
	case "BOTH_TEAMS_TO_SCORE":
	case "MATCH_ODDS":
		return odds, nil
	//case "MATCH_SHOTS":
	//case "MATCH_SHOTS_TARGET":
	case "OVER_UNDER_05":
		return parseMarketOdds("0.5", odds)
	case "OVER_UNDER_15":
		return parseMarketOdds("1.5", odds)
	case "OVER_UNDER_25":
		return parseMarketOdds("2.5", odds)
	case "OVER_UNDER_35":
		return parseMarketOdds("3.5", odds)
	case "OVER_UNDER_45":
		return parseMarketOdds("4.5", odds)
	case "OVER_UNDER_55_CORNR":
		return parseMarketOdds("5.5", odds)
	case "OVER_UNDER_85_CORNR":
		return parseMarketOdds("8.5", odds)
	case "OVER_UNDER_95_CORNR":
		return parseMarketOdds("9.5", odds)
	case "OVER_UNDER_105_CORNR":
		return parseMarketOdds("10.5", odds)
	case "OVER_UNDER_115_CORNR":
		return parseMarketOdds("11.5", odds)
	case "OVER_UNDER_125_CORNR":
		return parseMarketOdds("12.5", odds)
	case "OVER_UNDER_135_CORNR":
		return parseMarketOdds("13.5", odds)
	default:
		return nil, fmt.Errorf("market '%s' is not supported by exchange '%d'", market, exchangeID)
	}

	return nil, nil
}

func ParseExchange(exchangeID uint64, exchanges []sportmonks.MatchOdds) *sportmonks.BookmakerOdds {
	for _, market := range exchanges {
		for _, e := range market.BookmakerOdds() {
			if e.ID == int(exchangeID) {
				return &e
			}
		}
	}

	return nil
}

func parseMarketOdds(total string, odds []sportmonks.Odds) ([]sportmonks.Odds, error) {
	runners := []sportmonks.Odds{}

	for _, o := range odds {
		if o.Total == total {
			runners = append(runners, o)
		}
	}

	return runners, nil
}

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
