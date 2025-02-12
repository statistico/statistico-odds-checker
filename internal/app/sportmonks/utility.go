package sportmonks

import (
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-sportmonks-go-client"
	"strconv"
	"strings"
)

var marketIDs = map[string]int{
	"BOTH_TEAMS_TO_SCORE":     14,
	"MATCH_ODDS":              1,
	"PLAYER_TO_SCORE_ANYTIME": 90,
	"TEAM_CARDS":              281,
	"TEAM_CORNERS":            74,
	"TEAM_SHOTS":              285,
	"TEAM_SHOTS_ON_TARGET":    284,
}

func parseExchangeMarketOdds(exchangeId int, markets []sportmonks.Odds) []sportmonks.Odds {
	var odds []sportmonks.Odds

	for _, m := range markets {
		if m.BookmakerID == exchangeId {
			odds = append(odds, m)
		}
	}

	return odds
}

func convertOddsToRunners(odds []sportmonks.Odds, market string) ([]*exchange.Runner, error) {
	switch market {
	case "BOTH_TEAMS_TO_SCORE":
		return convertStandardOdds(odds)
	case "MATCH_ODDS":
		return convertStandardOdds(odds)
	case "PLAYER_TO_SCORE_ANYTIME":
		return convertPlayerToScore(odds, "ANYTIME")
	case "TEAM_CARDS":
		return convertOverUnderMarket(odds)
	case "TEAM_CORNERS":
		return convertOverUnderMarket(odds)
	case "TEAM_SHOTS":
		return convertOverUnderMarket(odds)
	case "TEAM_SHOTS_ON_TARGET":
		return convertOverUnderMarket(odds)
	default:
		return nil, fmt.Errorf("market %s is not supported", market)
	}
}

func convertStandardOdds(odds []sportmonks.Odds) ([]*exchange.Runner, error) {
	var runners []*exchange.Runner

	for _, o := range odds {
		price, err := strconv.ParseFloat(o.Value, 32)

		if err != nil {
			return nil, fmt.Errorf("value '%s' is not a valid floating point number", o.DP3)
		}

		runners = append(runners, &exchange.Runner{
			Name: strings.ToUpper(o.Label),
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

func convertOverUnderMarket(odds []sportmonks.Odds) ([]*exchange.Runner, error) {
	var runners []*exchange.Runner

	for _, o := range odds {
		price, err := strconv.ParseFloat(o.Value, 32)

		if err != nil {
			return nil, fmt.Errorf("value '%s' is not a valid floating point number", o.DP3)
		}

		team, err := parseTeam(o.Label)

		if err != nil {
			return nil, err
		}

		runners = append(runners, &exchange.Runner{
			Name: strings.ToUpper(fmt.Sprintf("%s - %s", team, *o.Total)),
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

func convertPlayerToScore(odds []sportmonks.Odds, label string) ([]*exchange.Runner, error) {
	var runners []*exchange.Runner

	for _, o := range odds {
		if strings.ToUpper(o.Label) != label {
			continue
		}

		price, err := strconv.ParseFloat(o.Value, 32)

		if err != nil {
			return nil, fmt.Errorf("value '%s' is not a valid floating point number", o.DP3)
		}

		runners = append(runners, &exchange.Runner{
			Name: *o.Name,
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

func parseTeam(label string) (string, error) {
	if label == "1" {
		return "HOME", nil
	}

	if label == "2" {
		return "AWAY", nil
	}

	return "", fmt.Errorf("invalid team label '%s'", label)
}
