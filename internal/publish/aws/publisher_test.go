package aws_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/statistico/statistico-odds-checker/internal/exchange"
	paws "github.com/statistico/statistico-odds-checker/internal/publish/aws"
	"github.com/statistico/statistico-odds-checker/internal/sport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestPublisher_PublishMarket(t *testing.T) {
	t.Run("publishes new message via AWS SNS client", func(t *testing.T) {
		t.Helper()

		client := new(paws.MockSnsClient)
		topic := "my-topic-arn"

		publisher := paws.NewPublisher(client, topic)

		input := mock.MatchedBy(func(i *sns.PublishInput) bool {
			ms := "{\"id\":\"1.23712\",\"eventId\":129817121,\"competitionId\":8,\"seasonId\":17420,\"sport\":\"football\"," +
				"\"date\":\"2019-01-14T11:00:00Z\",\"name\":\"1X2\",\"side\":\"BACK\"," +
				"\"exchange\":\"betfair\",\"runners\":[{\"id\":14571761,\"name\":\"Over 2.5 Goals\",\"sort\":1," +
				"\"prices\":[{\"price\":1.95,\"size\":1461}]}],\"timestamp\":1604430059}"

			assert.Equal(t, ms, *i.Message)
			assert.Equal(t, "my-topic-arn", *i.TopicArn)
			return true
		})

		client.On("Publish", input).Return(&sns.PublishOutput{}, nil)

		err := publisher.PublishMarket(eventMarket())

		if err != nil {
			t.Fatalf("Expected nil, got %s", err)
		}
	})

	t.Run("returns error if returned by AWS SNS client", func(t *testing.T) {
		t.Helper()

		client := new(paws.MockSnsClient)
		topic := "my-topic-arn"

		publisher := paws.NewPublisher(client, topic)

		input := mock.MatchedBy(func(i *sns.PublishInput) bool {
			ms := "{\"id\":\"1.23712\",\"eventId\":129817121,\"competitionId\":8,\"seasonId\":17420,\"sport\":\"football\"," +
				"\"date\":\"2019-01-14T11:00:00Z\",\"name\":\"1X2\",\"side\":\"BACK\"," +
				"\"exchange\":\"betfair\",\"runners\":[{\"id\":14571761,\"name\":\"Over 2.5 Goals\",\"sort\":1,"+
				"\"prices\":[{\"price\":1.95,\"size\":1461}]}],\"timestamp\":1604430059}"

			assert.Equal(t, ms, *i.Message)
			assert.Equal(t, "my-topic-arn", *i.TopicArn)
			return true
		})

		client.On("Publish", input).Return(&sns.PublishOutput{}, errors.New("error"))

		err := publisher.PublishMarket(eventMarket())

		if err == nil {
			t.Fatal("Expected error, got nil")
		}
	})
}

func eventMarket() *sport.EventMarket {
	return &sport.EventMarket{
		ID:             "1.23712",
		EventID:        129817121,
		CompetitionID:  8,
		SeasonID:       17420,
		Sport:          "football",
		EventDate:      "2019-01-14T11:00:00Z",
		MarketName:     "1X2",
		Side:           "BACK",
		Exchange:       "betfair",
		Runners: []*exchange.Runner{
			{
				ID: 14571761,
				Name: "Over 2.5 Goals",
				Sort: 1,
				Prices: []exchange.PriceSize{
					{
						Price: 1.95,
						Size: 1461.00,
					},
				},
			},
		},
		Timestamp:      1604430059,
	}
}