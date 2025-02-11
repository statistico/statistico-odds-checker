package aws

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/statistico/statistico-odds-checker/internal/app/publish"
	"github.com/statistico/statistico-odds-checker/internal/app/stream"
)

type Publisher struct {
	client   snsiface.SNSAPI
	topicArn string
}

func (p *Publisher) PublishMarket(m *stream.EventMarket) error {
	js, err := json.Marshal(m)

	if err != nil {
		return err
	}

	input := sns.PublishInput{
		Message:  aws.String(string(js)),
		TopicArn: aws.String(p.topicArn),
	}

	mess, err := p.client.Publish(&input)

	if err != nil {
		return err
	}

	fmt.Printf("Published message %s", *mess.MessageId)

	return nil
}

func NewPublisher(c snsiface.SNSAPI, t string) publish.Publisher {
	return &Publisher{
		client:   c,
		topicArn: t,
	}
}
