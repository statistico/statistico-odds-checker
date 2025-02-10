package sportmonks

import (
	"fmt"
	"github.com/statistico/statistico-sportmonks-go-client"
)

var marketIDs = map[string]int{
	"BOTH_TEAMS_TO_SCORE": 14,
	"MATCH_ODDS":          1,
}

func parseMarketRunners(market string, exchangeID int, odds []sportmonks.Odds) ([]sportmonks.Odds, error) {
	switch market {
	case "BOTH_TEAMS_TO_SCORE":
		return odds, nil
	case "MATCH_ODDS":
		return odds, nil
	default:
		return nil, fmt.Errorf("market '%s' is not supported by exchange '%d'", market, exchangeID)
	}
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
