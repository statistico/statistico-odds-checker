package betfair_test

import (
	"bytes"
	"context"
	betfair "github.com/statistico/statistico-betfair-go-client"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	bf "github.com/statistico/statistico-odds-checker/internal/app/exchange/betfair"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestMarketRequester_Fetch(t *testing.T) {
	url := betfair.BaseURLs{
		Accounts: "https://mock.com",
		Betting:  "https://mock.com",
		Login:    "https://mock.com/login",
	}

	e := exchange.Event{
		Name:   "West Ham United v Manchester City",
		Date:   time.Date(2020, 10, 24, 12, 30, 00, 0, time.UTC),
		Market: "OVER_UNDER_25",
	}

	t.Run("calls betfair service using client and returns a exchange market struct", func(t *testing.T) {
		t.Helper()

		tc := NewTestClient(func(req *http.Request) *http.Response {
			if req.URL.Path == "/login" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(loginResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listEvents/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(eventsResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listMarketCatalogue/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(marketCatalogueResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listRunnerBook/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(runnersResponse)),
					Header:     make(http.Header),
				}
			}

			return nil
		})

		client := betfair.Client{HTTPClient: tc, BaseURLs: url}

		factory := bf.NewMarketFactory(client)

		market, err := factory.CreateMarket(context.Background(), &e)

		if err != nil {
			t.Fatalf("Error: %s", err.Error())
		}

		a := assert.New(t)

		a.Equal("1.173887003", market.ID)
		a.Equal(uint64(47972), market.Runners[0].ID)
		a.Equal("Under 2.5 Goals", market.Runners[0].Name)
		a.Equal(float32(2.96), market.Runners[0].BackPrices[0].Price)
		a.Equal(float32(152.84), market.Runners[0].BackPrices[0].Size)
		a.Equal(float32(2.9), market.Runners[0].BackPrices[1].Price)
		a.Equal(float32(34.5), market.Runners[0].BackPrices[1].Size)
		a.Equal(float32(2.88), market.Runners[0].BackPrices[2].Price)
		a.Equal(float32(91.04), market.Runners[0].BackPrices[2].Size)
		a.Equal(float32(2.96), market.Runners[0].LayPrices[0].Price)
		a.Equal(float32(152.84), market.Runners[0].LayPrices[0].Size)
		a.Equal(float32(2.9), market.Runners[0].LayPrices[1].Price)
		a.Equal(float32(34.5), market.Runners[0].LayPrices[1].Size)
		a.Equal(float32(2.88), market.Runners[0].LayPrices[2].Price)
		a.Equal(float32(91.04), market.Runners[0].LayPrices[2].Size)

		a.Equal(uint64(47973), market.Runners[1].ID)
		a.Equal("Over 2.5 Goals", market.Runners[1].Name)
		a.Equal(float32(2.96), market.Runners[0].BackPrices[0].Price)
		a.Equal(float32(152.84), market.Runners[0].BackPrices[0].Size)
		a.Equal(float32(2.9), market.Runners[0].BackPrices[1].Price)
		a.Equal(float32(34.5), market.Runners[0].BackPrices[1].Size)
		a.Equal(float32(2.88), market.Runners[0].BackPrices[2].Price)
		a.Equal(float32(91.04), market.Runners[0].BackPrices[2].Size)
		a.Equal(float32(2.96), market.Runners[0].LayPrices[0].Price)
		a.Equal(float32(152.84), market.Runners[0].LayPrices[0].Size)
		a.Equal(float32(2.9), market.Runners[0].LayPrices[1].Price)
		a.Equal(float32(34.5), market.Runners[0].LayPrices[1].Size)
		a.Equal(float32(2.88), market.Runners[0].LayPrices[2].Price)
		a.Equal(float32(91.04), market.Runners[0].LayPrices[2].Size)
	})

	t.Run("returns no event error if events response is empty", func(t *testing.T) {
		t.Helper()

		tc := NewTestClient(func(req *http.Request) *http.Response {
			if req.URL.Path == "/login" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(loginResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listEvents/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`[]`)),
					Header:     make(http.Header),
				}
			}

			return nil
		})

		client := betfair.Client{HTTPClient: tc, BaseURLs: url}

		factory := bf.NewMarketFactory(client)

		_, err := factory.CreateMarket(context.Background(), &e)

		if err == nil {
			t.Fatalf("Expected error got nil")
		}

		assert.Equal(t, "No event returned for: West Ham v Man City", err.Error())
	})

	t.Run("returns multiple market error if more than one market catalogue returned", func(t *testing.T) {
		t.Helper()

		tc := NewTestClient(func(req *http.Request) *http.Response {
			if req.URL.Path == "/login" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(loginResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listEvents/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(eventsResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listMarketCatalogue/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`[{}, {}]`)),
					Header:     make(http.Header),
				}
			}

			return nil
		})

		client := betfair.Client{HTTPClient: tc, BaseURLs: url}

		factory := bf.NewMarketFactory(client)

		_, err := factory.CreateMarket(context.Background(), &e)

		if err == nil {
			t.Fatalf("Expected error got nil")
		}

		assert.Equal(t, "Multiple markets returned for event: 30066485", err.Error())
	})

	t.Run("returns no market error if no runners returned for a given market", func(t *testing.T) {
		t.Helper()

		tc := NewTestClient(func(req *http.Request) *http.Response {
			if req.URL.Path == "/login" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(loginResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listEvents/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(eventsResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listMarketCatalogue/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(marketCatalogueResponse)),
					Header:     make(http.Header),
				}
			}

			if req.URL.Path == "/listRunnerBook/" {
				return &http.Response{
					StatusCode: 200,
					Body:       ioutil.NopCloser(bytes.NewBufferString(`[{}, {}]`)),
					Header:     make(http.Header),
				}
			}

			return nil
		})

		client := betfair.Client{HTTPClient: tc, BaseURLs: url}

		factory := bf.NewMarketFactory(client)

		_, err := factory.CreateMarket(context.Background(), &e)

		if err == nil {
			t.Fatalf("Expected error got nil")
		}

		assert.Equal(t, "Multiple selections returned for market 1.173887003 and selection 47972", err.Error())
	})
}

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}

var loginResponse = `{
	"token": "token123",
	"product": "soccer",
	"status": "SUCCESS",
	"error": ""
}`

var eventsResponse = `[
  {
    "event": {
      "id": "30066485",
      "name": "Cagliari v Crotone",
      "countryCode": "IT",
      "timezone": "GMT",
      "openDate": "2020-10-25T11:30:00.000Z"
    },
    "marketCount": 7
  }
]`

var marketCatalogueResponse = `[
  {
    "marketId": "1.173887003",
    "marketName": "Over/Under 2.5 Goals",
    "totalMatched": 0.0,
    "runners": [
      {
        "selectionId": 47972,
        "runnerName": "Under 2.5 Goals",
        "handicap": 0.0,
        "sortPriority": 1,
        "metadata": {
          "runnerId": "47972"
        }
      },
      {
        "selectionId": 47973,
        "runnerName": "Over 2.5 Goals",
        "handicap": 0.0,
        "sortPriority": 2,
        "metadata": {
          "runnerId": "47973"
        }
      }
    ]
  }
]`

var runnersResponse = `[
  {
    "marketId": "1.173887003",
    "runners": [
      {
        "selectionId": 47972,
        "handicap": 0.0,
        "status": "ACTIVE",
        "totalMatched": 0.0,
        "ex": {
          "availableToBack": [
            {
              "price": 2.96,
              "size": 152.84
            },
            {
              "price": 2.9,
              "size": 34.5
            },
            {
              "price": 2.88,
              "size": 91.04
            }
          ],
          "availableToLay": [
            {
              "price": 2.96,
              "size": 152.84
            },
            {
              "price": 2.9,
              "size": 34.5
            },
            {
              "price": 2.88,
              "size": 91.04
            }
          ],
          "tradedVolume": []
        }
      }
    ]
  }
]`
