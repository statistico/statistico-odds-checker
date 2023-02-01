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

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 1, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		underOdds := spClient.FlexFloat(1.91)
		overOdds := spClient.FlexFloat(2.65)

		expectedOdds := []spClient.Odds{
			{
				Value:            &underOdds,
				Handicap:         nil,
				Total:            "2.5",
				Label:            "Under",
				Probability:      "52.36%",
				Dp3:              "1.910",
				American:         -110,
				Fractional:       nil,
				Winning:          nil,
				Stop:             true,
				BookmakerEventID: nil,
				LastUpdate: spClient.DateTime{
					Date:         "2019-10-05 13:01:00.227530",
					TimezoneType: 3,
					Timezone:     "UTC",
				},
			},
			{
				Value:            &overOdds,
				Handicap:         nil,
				Total:            "2.5",
				Label:            "Over",
				Probability:      "52.36%",
				Dp3:              "1.910",
				American:         -110,
				Fractional:       nil,
				Winning:          nil,
				Stop:             true,
				BookmakerEventID: nil,
				LastUpdate: spClient.DateTime{
					Date:         "2019-10-05 13:01:00.227530",
					TimezoneType: 2,
					Timezone:     "UTC",
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

		_, err := parser.ParseMarketOdds(context.Background(), 152, 1, "ASIAN_HANDICAP")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error handling market for exchange '1': market 'ASIAN_HANDICAP' is not supported", err.Error())
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

		_, err := parser.ParseMarketOdds(context.Background(), 152, 1, "OVER_UNDER_25")

		if err == nil {
			t.Fatal("Expected error, got nil")
		}

		assert.Equal(t, "error fetching markets for exchange '1': Request failed with message: The requested endpoint does not exist!, code: 404", err.Error())
	})

	t.Run("returns nil if no markets are returned for event", func(t *testing.T) {
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

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 1, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Nil(odds)
	})

	t.Run("returns nil if no market provided for exchange", func(t *testing.T) {
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

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 111, "OVER_UNDER_25")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Nil(odds)
	})

	t.Run("returns nil if no odds provided for exchange market", func(t *testing.T) {
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

		odds, err := parser.ParseMarketOdds(context.Background(), 152, 123, "OVER_UNDER_45")

		if err != nil {
			t.Fatalf("Expected nil, got %s", err.Error())
		}

		a := assert.New(t)

		a.Nil(odds)
	})
}

var overUnderGoalsOddsResponse = `{
	"data": [
		{
			"id": 38,
			"name": "Over\\/Under",
			"suspended": false,
			"bookmaker": {
				"data": [
					{
						"id": 1,
            			"name": "10Bet",
						"odds": {
							"data": [
								{
									"value": "1.91",
									"handicap": null,
									"total": "2.5",
									"label": "Under",
									"probability": "52.36%",
									"dp3": "1.910",
									"american": -110,
									"factional": null,
									"winning": null,
									"stop": true,
									"bookmaker_event_id": null,
									"last_update": {
										"date": "2019-10-05 13:01:00.227530",
										"timezone_type": 3,
										"timezone": "UTC"
									}
								},
								{
									"value": "2.65",
									"handicap": null,
									"total": "2.5",
									"label": "Over",
									"probability": "52.36%",
									"dp3": "1.910",
									"american": -110,
									"factional": null,
									"winning": null,
									"stop": true,
									"bookmaker_event_id": null,
									"last_update": {
										"date": "2019-10-05 13:01:00.227530",
										"timezone_type": 2,
										"timezone": "UTC"
									}
								}
							]
						}
					}
				]
			}
		}
	]
}`

var EmptyOddsResponse = `{
	"data": []
}`

var errorResponse = `{
	"error": {
		"message": "The requested endpoint does not exist!",
		"code": 404
	}
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
