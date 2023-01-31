package sportmonks

import (
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
)

func parseMarketId(market string) (int, error) {
	for id, m := range markets {
		for _, mk := range m {
			if mk == market {
				return id, nil
			}
		}
	}

	return 0, fmt.Errorf("market '%s' is not supported", market)
}

func parseMarketRunners(market string, exchangeID int, odds []sportmonks.Odds) ([]sportmonks.Odds, error) {
	switch market {
	case "BOTH_TEAMS_TO_SCORE":
		return odds, nil
	case "MATCH_ODDS":
		return odds, nil
	case "OVER_UNDER_05":
		return parseMarketOdds([]string{"0.5"}, odds)
	case "OVER_UNDER_15":
		return parseMarketOdds([]string{"1.5"}, odds)
	case "OVER_UNDER_25":
		return parseMarketOdds([]string{"2.25", "2.5"}, odds)
	case "OVER_UNDER_35":
		return parseMarketOdds([]string{"3.5"}, odds)
	case "OVER_UNDER_45":
		return parseMarketOdds([]string{"4.5"}, odds)
	case "OVER_UNDER_55_CORNR":
		return parseMarketOdds([]string{"5.5"}, odds)
	case "OVER_UNDER_85_CORNR":
		return parseMarketOdds([]string{"8.5"}, odds)
	case "OVER_UNDER_95_CORNR":
		return parseMarketOdds([]string{"9.5"}, odds)
	case "OVER_UNDER_105_CORNR":
		return parseMarketOdds([]string{"10.5"}, odds)
	case "OVER_UNDER_115_CORNR":
		return parseMarketOdds([]string{"11.5"}, odds)
	case "OVER_UNDER_125_CORNR":
		return parseMarketOdds([]string{"12.5"}, odds)
	case "OVER_UNDER_135_CORNR":
		return parseMarketOdds([]string{"13.5"}, odds)
	default:
		return nil, fmt.Errorf("market '%s' is not supported by exchange '%d'", market, exchangeID)
	}
}

func parseMarketOdds(total []string, odds []sportmonks.Odds) ([]sportmonks.Odds, error) {
	runners := []sportmonks.Odds{}

	for _, o := range odds {
		for _, t := range total {
			if o.Total == t {
				runners = append(runners, o)
			}
		}
	}

	return runners, nil
}

func parseExchangeOdds(exchangeID int, exchanges []sportmonks.MatchOdds) []sportmonks.Odds {
	for _, market := range exchanges {
		for _, e := range market.BookmakerOdds() {
			if e.ID == exchangeID {
				return e.Odds()
			}
		}
	}

	return nil
}
