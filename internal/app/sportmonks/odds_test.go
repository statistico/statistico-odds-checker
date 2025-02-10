package sportmonks_test

import (
	"bytes"
	"context"
	"github.com/statistico/statistico-odds-checker/internal/app/sportmonks"
	spClient "github.com/statistico/statistico-sportmonks-go-client"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestOddsParser_ParseMarketOdds(t *testing.T) {
	t.Run("calls sportmonks api and parses odds for event, exchange and market", func(t *testing.T) {
		t.Helper()

		server := httpClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(overUnderGoalsOddsResponse)),
			}, nil
		})

		client := spClient.HTTPClient{
			HTTPClient: server,
			BaseURL:    "https://example.com",
			Key:        "my-key",
		}

		parser := sportmonks.NewOddsParser(&client)

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 2, "MATCH_ODDS")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		labelOne := "1"
		labelTwo := "2"

		created, _ := time.Parse(time.RFC3339, "2025-01-26T14:57:19.000000Z")

		expectedOdds := []spClient.Odds{
			{
				ID:                    151577019200,
				FixtureID:             19155301,
				MarketID:              1,
				BookmakerID:           2,
				Label:                 "Home",
				Value:                 "3.75",
				Name:                  &labelOne,
				SortOrder:             nil,
				MarketDescription:     "Full Time Result",
				Probability:           "26.67%",
				DP3:                   "1.750",
				Fractional:            "15/4",
				American:              "275",
				Winning:               false,
				Stopped:               false,
				Total:                 nil,
				Handicap:              nil,
				Participants:          nil,
				CreatedAt:             created,
				OriginalLabel:         nil,
				LatestBookmakerUpdate: "2025-02-10 14:10:51",
				Bookmaker: &spClient.Bookmaker{
					ID:       2,
					LegacyID: 2,
					Name:     "bet365",
				},
				Market: &spClient.Market{
					ID:                     1,
					LegacyID:               1,
					Name:                   "Fulltime Result",
					DeveloperName:          "FULLTIME_RESULT",
					HasWinningCalculations: false,
				},
			},
			{
				ID:                    151577019200,
				FixtureID:             19155301,
				MarketID:              1,
				BookmakerID:           2,
				Label:                 "Away",
				Value:                 "3.75",
				Name:                  &labelTwo,
				SortOrder:             nil,
				MarketDescription:     "Full Time Result",
				Probability:           "26.67%",
				DP3:                   "3.750",
				Fractional:            "15/4",
				American:              "275",
				Winning:               false,
				Stopped:               false,
				Total:                 nil,
				Handicap:              nil,
				Participants:          nil,
				CreatedAt:             created,
				OriginalLabel:         nil,
				LatestBookmakerUpdate: "2025-02-10 14:10:51",
				Bookmaker: &spClient.Bookmaker{
					ID:       2,
					LegacyID: 2,
					Name:     "bet365",
				},
				Market: &spClient.Market{
					ID:                     1,
					LegacyID:               1,
					Name:                   "Fulltime Result",
					DeveloperName:          "FULLTIME_RESULT",
					HasWinningCalculations: false,
				},
			},
		}

		a := assert.New(t)

		a.Equal(2, len(odds))
		a.Equal(expectedOdds, odds)
	})

	t.Run("returns an error if unable to parse market id", func(t *testing.T) {
		t.Helper()

		client := spClient.HTTPClient{
			HTTPClient: nil,
			BaseURL:    "https://example.com",
			Key:        "my-key",
		}

		parser := sportmonks.NewOddsParser(&client)

		_, err := parser.ParseMarketOdds(context.Background(), 152, 2, "ASIAN_HANDICAP")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error handling market for exchange '2': market 'ASIAN_HANDICAP' is not supported", err.Error())
	})

	t.Run("return an error if error returned by sportmonks client", func(t *testing.T) {
		t.Helper()

		server := httpClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewBufferString(errorResponse)),
			}, nil
		})

		client := spClient.HTTPClient{
			HTTPClient: server,
			BaseURL:    "https://example.com",
			Key:        "my-key",
		}

		parser := sportmonks.NewOddsParser(&client)

		_, err := parser.ParseMarketOdds(context.Background(), 152, 2, "MATCH_ODDS")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error fetching markets for exchange '2': Request failed with the message: The requested endpoint does not exist!", err.Error())
	})

	t.Run("returns an empty slice of struct if no markets are returned for event", func(t *testing.T) {
		server := httpClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(EmptyOddsResponse)),
			}, nil
		})

		client := spClient.HTTPClient{
			HTTPClient: server,
			BaseURL:    "https://example.com",
			Key:        "my-key",
		}

		parser := sportmonks.NewOddsParser(&client)

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 1, "MATCH_ODDS")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(0, len(odds))
	})

	t.Run("returns an empty slice of struct if no market provided for exchange", func(t *testing.T) {
		server := httpClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(overUnderGoalsOddsResponse)),
			}, nil
		})

		client := spClient.HTTPClient{
			HTTPClient: server,
			BaseURL:    "https://example.com",
			Key:        "my-key",
		}

		parser := sportmonks.NewOddsParser(&client)

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 111, "MATCH_ODDS")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(0, len(odds))
	})

	t.Run("returns an empty slice of struct if no odds provided for exchange market", func(t *testing.T) {
		server := httpClient(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBufferString(overUnderGoalsOddsResponse)),
			}, nil
		})

		client := spClient.HTTPClient{
			HTTPClient: server,
			BaseURL:    "https://example.com",
			Key:        "my-key",
		}

		parser := sportmonks.NewOddsParser(&client)

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 123, "MATCH_ODDS")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Equal(0, len(odds))
	})
}

var overUnderGoalsOddsResponse = `{
	"data": [
		{
			"id": 151577019200,
			"fixture_id": 19155301,
			"market_id": 1,
			"bookmaker_id": 2,
			"label": "Home",
			"value": "3.75",
			"name": "1",
			"sort_order": null,
			"market_description": "Full Time Result",
			"probability": "26.67%",
			"dp3": "1.750",
			"fractional": "15\/4",
			"american": "275",
			"winning": false,
			"stopped": false,
			"total": null,
			"handicap": null,
			"participants": null,
			"created_at": "2025-01-26T14:57:19.000000Z",
			"original_label": null,
			"latest_bookmaker_update": "2025-02-10 14:10:51",
			"bookmaker": {
				"id": 2,
				"legacy_id": 2,
				"name": "bet365"
			},
			"market": {
				"id": 1,
				"legacy_id": 1,
				"name": "Fulltime Result",
				"developer_name": "FULLTIME_RESULT",
				"has_winning_calculations": false
			}
		},
		{
			"id": 151577019200,
			"fixture_id": 19155301,
			"market_id": 1,
			"bookmaker_id": 2,
			"label": "Away",
			"value": "3.75",
			"name": "2",
			"sort_order": null,
			"market_description": "Full Time Result",
			"probability": "26.67%",
			"dp3": "3.750",
			"fractional": "15\/4",
			"american": "275",
			"winning": false,
			"stopped": false,
			"total": null,
			"handicap": null,
			"participants": null,
			"created_at": "2025-01-26T14:57:19.000000Z",
			"original_label": null,
			"latest_bookmaker_update": "2025-02-10 14:10:51",
			"bookmaker": {
				"id": 2,
				"legacy_id": 2,
				"name": "bet365"
			},
			"market": {
				"id": 1,
				"legacy_id": 1,
				"name": "Fulltime Result",
				"developer_name": "FULLTIME_RESULT",
				"has_winning_calculations": false
			}
		}
	]
}`

var EmptyOddsResponse = `{
	"data": []
}`

var errorResponse = `{
	"message": "The requested endpoint does not exist!",
	"code": 404
}`

func httpClient(fn roundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

type roundTripFunc func(req *http.Request) (*http.Response, error)

func (r roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return r(req)
}
