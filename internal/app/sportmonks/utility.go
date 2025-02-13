package sportmonks

import (
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-sportmonks-go-client"
	"strconv"
	"strings"
)

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
	case "MATCH_CORNERS":
		return convertMatchOverUnderMarket(odds)
	case "MATCH_GOALS":
		return convertMatchOverUnderMarket(odds)
	case "MATCH_RESULT":
		return convertStandardOdds(odds)
	case "MATCH_SHOTS_ON_TARGET":
		return convertMatchOverUnderMarket(odds)
	case "MATCH_SHOTS_TOTAL":
		return convertMatchOverUnderMarket(odds)
	case "PLAYER_CARDS":
		return convertPlayerCards(odds, "BOOKED")
	case "PLAYER_SHOTS_ON_TARGET":
		return convertPlayerOverUnder(odds, "Player Shots On Target Over\\/Under")
	case "PLAYER_SHOTS_TOTAL":
		return convertPlayerOverUnder(odds, strings.ToUpper("Player Shots Over\\/Under"))
	case "PLAYER_TACKLES":
		return convertPlayerOverUnder(odds, "Player Tackles")
	case "PLAYER_TO_SCORE_ANYTIME":
		return convertPlayerToScore(odds, "ANYTIME")
	case "TEAM_CARDS":
		return convertTeamOverUnderMarket(odds)
	case "TEAM_CORNERS":
		return convertTeamOverUnderMarket(odds)
	case "TEAM_GOALS":
		return convertTeamOverUnderMarket(odds)
	case "TEAM_SHOTS":
		return convertTeamOverUnderMarket(odds)
	case "TEAM_SHOTS_ON_TARGET":
		return convertTeamOverUnderMarket(odds)
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
			ID:   strconv.Itoa(o.ID),
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

func convertMatchOverUnderMarket(odds []sportmonks.Odds) ([]*exchange.Runner, error) {
	var runners []*exchange.Runner

	for _, o := range odds {
		price, err := strconv.ParseFloat(o.Value, 32)

		if err != nil {
			return nil, fmt.Errorf("price '%s' is not a valid floating point number", o.Value)
		}

		runners = append(runners, &exchange.Runner{
			ID:   strconv.Itoa(o.ID),
			Name: fmt.Sprintf("%s %s", strings.ToUpper(o.Label), *o.Total),
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

func convertTeamOverUnderMarket(odds []sportmonks.Odds) ([]*exchange.Runner, error) {
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

		var label string

		if o.Total != nil {
			label = strings.ToUpper(*o.Total)
		}

		runners = append(runners, &exchange.Runner{
			ID:    strconv.Itoa(o.ID),
			Name:  team,
			Label: &label,
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

func convertPlayerCards(odds []sportmonks.Odds, label string) ([]*exchange.Runner, error) {
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
			ID:    strconv.Itoa(o.ID),
			Name:  *o.Name,
			Label: &label,
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
			ID:    strconv.Itoa(o.ID),
			Name:  *o.Name,
			Label: &label,
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

func convertPlayerOverUnder(odds []sportmonks.Odds, description string) ([]*exchange.Runner, error) {
	var runners []*exchange.Runner

	for _, o := range odds {
		if strings.ToUpper(o.MarketDescription) != description {
			continue
		}

		price, err := strconv.ParseFloat(o.Value, 32)

		if err != nil {
			return nil, fmt.Errorf("price '%s' is not a valid floating point number", o.DP3)
		}

		label := fmt.Sprintf("%s %s", strings.ToUpper(o.Label), *o.Total)

		runners = append(runners, &exchange.Runner{
			ID:    strconv.Itoa(o.ID),
			Name:  *o.Name,
			Label: &label,
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
