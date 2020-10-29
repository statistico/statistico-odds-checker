package betfair

import (
	"context"
	"fmt"
	betfair "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	"strings"
	"time"
)

type MarketRequester struct {
	betfairClient betfair.Client
}

func (m *MarketRequester) Fetch(ctx context.Context, q *exchange.Query) (*exchange.Market, error) {
	event, err := m.getEvent(ctx, q)

	if err != nil {
		return nil, err
	}

	return m.parseMarket(ctx, buildMarketCatalogueRequest(event.ID, q))
}

func (m *MarketRequester) getEvent(ctx context.Context, q *exchange.Query) (*betfair.Event, error) {
	req := buildEventsRequest(q)

	events, err := m.betfairClient.ListEvents(ctx, req)

	if err != nil {
		return nil, &clientError{"list events", err}
	}

	if len(events) == 0 {
		return nil, &noEventError{req.Filter.TextQuery}
	}

	return &events[0].Event, nil
}

func (m *MarketRequester) parseMarket(ctx context.Context, req betfair.ListMarketCatalogueRequest) (*exchange.Market, error) {
	catalogue, err := m.betfairClient.ListMarketCatalogue(ctx, req)

	if err != nil {
		return nil, &clientError{"market catalogue", err}
	}

	if len(catalogue) > 1 {
		return nil, &multipleEventMarketsError{strings.Join(req.Filter.EventIDs,  ",")}
	}

	market := exchange.Market{ID: catalogue[0].MarketID}

	for _, runner := range catalogue[0].Runners {
		prices, err := m.parseRunnerPrices(ctx, buildRunnerBookRequest(market.ID, runner.SelectionID))

		if err != nil {
			return nil, err
		}

		r := exchange.Runner{
			ID:     runner.SelectionID,
			Name:   runner.RunnerName,
			Prices: prices,
		}

		market.Runners = append(market.Runners, r)
	}

	return &market, nil
}

func (m *MarketRequester) parseRunnerPrices(ctx context.Context, req betfair.ListRunnerBookRequest) ([]exchange.PriceSize, error) {
	response, err := m.betfairClient.ListRunnerBook(ctx, req)

	if err != nil {
		return nil, &clientError{"list runner book", err}
	}

	if len(response) > 1 {
		return nil, &multipleMarketSelectionError{req.MarketID, req.SelectionID}
	}

	prices := []exchange.PriceSize{}

	for _, runner := range response[0].Runners {
		for _, price := range runner.EX.AvailableToBack {
			ps := exchange.PriceSize{
				Price: price.Price,
				Size:  price.Size,
			}

			prices = append(prices, ps)
		}
	}

	return prices, nil
}

func buildEventsRequest(q *exchange.Query) betfair.ListEventsRequest {
	from := q.Date.AddDate(0, 0, -1)
	to := q.Date.AddDate(0, 0, 1)

	text := q.Event

	if q.Sport == "football" {
		split := strings.Split(q.Event, " v ")
		text = fmt.Sprintf("%s v %s", teams[split[0]], teams[split[1]])
	}

	dates := betfair.TimeRange{
		From: from.Format(time.RFC3339),
		To: to.Format(time.RFC3339),
	}

	filter := betfair.MarketFilter{
		TextQuery:          text,
		MarketStartTime:    dates,
	}

	return betfair.ListEventsRequest{Filter: filter}
}

func buildMarketCatalogueRequest(eventID string, q *exchange.Query) betfair.ListMarketCatalogueRequest {
	eventIDs := []string{eventID}
	codes := []string{q.Market}
	projection := []string{"RUNNER_METADATA"}

	filter := betfair.MarketFilter{
		EventIDs: eventIDs,
		MarketTypeCodes: codes,
	}

	return betfair.ListMarketCatalogueRequest{
		Filter: filter,
		MarketProjection: projection,
		MaxResults: 1,
	}
}

func buildRunnerBookRequest(marketID string, selectionID uint64) betfair.ListRunnerBookRequest {
	projection := betfair.PriceProjection{PriceData: []string{"EX_BEST_OFFERS"}}

	return betfair.ListRunnerBookRequest{
		MarketID:         marketID,
		SelectionID:      selectionID,
		PriceProjection:  projection,
	}
}

func NewMarketRequester(c betfair.Client) MarketRequester {
	return MarketRequester{betfairClient: c}
}