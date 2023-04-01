package betfair

import (
	"context"
	"fmt"
	betfair "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"strings"
	"time"
)

type marketFactory struct {
	betfairClient betfair.Client
}

func (*marketFactory) Exchange() string {
	return "BETFAIR"
}

func (m *marketFactory) CreateMarket(ctx context.Context, q *exchange.Event) (*exchange.Market, error) {
	event, err := m.getEvent(ctx, q)

	if err != nil {
		return nil, err
	}

	req := buildMarketCatalogueRequest(event.ID, q)

	return m.parseMarket(ctx, req, q)
}

func (m *marketFactory) getEvent(ctx context.Context, q *exchange.Event) (*betfair.Event, error) {
	req := buildEventsRequest(q)

	events, err := m.betfairClient.ListEvents(ctx, req)

	if err != nil {
		return nil, &exchange.ClientError{Context: "list events", E: err}
	}

	if len(events) == 0 {
		return nil, &exchange.NoEventError{Event: req.Filter.TextQuery}
	}

	return &events[0].Event, nil
}

func (m *marketFactory) parseMarket(ctx context.Context, req betfair.ListMarketCatalogueRequest, q *exchange.Event) (*exchange.Market, error) {
	catalogue, err := m.betfairClient.ListMarketCatalogue(ctx, req)

	if err != nil {
		return nil, &exchange.ClientError{Context: "market catalogue", E: err}
	}

	if len(catalogue) == 0 {
		return nil, &exchange.NoEventMarketError{
			Exchange: "BETFAIR",
			Market:   q.Market,
			EventID:  q.ID,
		}
	}

	if len(catalogue) > 1 {
		return nil, &exchange.MultipleEventMarketsError{EventID: strings.Join(req.Filter.EventIDs, ",")}
	}

	market := exchange.Market{
		ID:       catalogue[0].MarketID,
		EventID:  q.ID,
		Name:     q.Market,
		Exchange: "BETFAIR",
	}

	for _, runner := range catalogue[0].Runners {
		back, lay, err := m.parseRunnerPrices(ctx, buildRunnerBookRequest(market.ID, runner.SelectionID))

		if err != nil {
			return nil, err
		}

		r := &exchange.Runner{
			ID:         runner.SelectionID,
			Name:       parseRunnerName(&runner, q.Market),
			BackPrices: back,
			LayPrices:  lay,
		}

		market.Runners = append(market.Runners, r)
	}

	return &market, nil
}

func (m *marketFactory) parseRunnerPrices(ctx context.Context, req betfair.ListRunnerBookRequest) ([]exchange.PriceSize, []exchange.PriceSize, error) {
	response, err := m.betfairClient.ListRunnerBook(ctx, req)

	if err != nil {
		return nil, nil, &exchange.ClientError{Context: "list runner book", E: err}
	}

	if len(response) != 1 {
		return nil, nil, &exchange.MultipleMarketSelectionError{EventID: req.MarketID, SelectionID: req.SelectionID}
	}

	back := []exchange.PriceSize{}
	lay := []exchange.PriceSize{}

	for _, runner := range response[0].Runners {
		for _, price := range runner.EX.AvailableToBack {
			ps := exchange.PriceSize{
				Price: price.Price,
				Size:  price.Size,
			}

			back = append(back, ps)
		}

		for _, price := range runner.EX.AvailableToLay {
			ps := exchange.PriceSize{
				Price: price.Price,
				Size:  price.Size,
			}

			lay = append(lay, ps)
		}
	}

	return back, lay, nil
}

func buildEventsRequest(e *exchange.Event) betfair.ListEventsRequest {
	from := e.Date.AddDate(0, 0, -1)
	to := e.Date.AddDate(0, 0, 1)

	split := strings.Split(e.Name, " v ")

	dates := betfair.TimeRange{
		From: from.Format(time.RFC3339),
		To:   to.Format(time.RFC3339),
	}

	filter := betfair.MarketFilter{
		TextQuery:       fmt.Sprintf("%s v %s", teams[split[0]], teams[split[1]]),
		MarketStartTime: dates,
	}

	return betfair.ListEventsRequest{Filter: filter}
}

func buildMarketCatalogueRequest(eventID string, e *exchange.Event) betfair.ListMarketCatalogueRequest {
	eventIDs := []string{eventID}
	codes := []string{e.Market}
	projection := []string{"RUNNER_METADATA"}

	filter := betfair.MarketFilter{
		EventIDs:        eventIDs,
		MarketTypeCodes: codes,
	}

	return betfair.ListMarketCatalogueRequest{
		Filter:           filter,
		MarketProjection: projection,
		MaxResults:       1,
	}
}

func buildRunnerBookRequest(marketID string, selectionID uint64) betfair.ListRunnerBookRequest {
	projection := betfair.PriceProjection{PriceData: []string{"EX_BEST_OFFERS"}}

	return betfair.ListRunnerBookRequest{
		MarketID:        marketID,
		SelectionID:     selectionID,
		PriceProjection: projection,
	}
}

func NewMarketFactory(c betfair.Client) exchange.MarketFactory {
	return &marketFactory{betfairClient: c}
}
