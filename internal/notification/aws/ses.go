package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/statistico/statistico-odds-checker/internal/notification"
)

var charset = "UTF-8"

type sesNotifier struct {
	client sesiface.SESAPI
}

func (s *sesNotifier) SendMessage(m *notification.Message) *notification.Result {
	email := buildEmailInput(m)

	result, err := s.client.SendEmail(email)

	if err != nil {
		return &notification.Result{
			Status:  "FAIL",
			Message: err.Error(),
		}
	}

	return &notification.Result{
		Status:  "SUCCESS",
		Message: *result.MessageId,
	}
}

func buildEmailInput(m *notification.Message) *ses.SendEmailInput {
	subject := &ses.Content{
		Charset: aws.String(charset),
		Data:    aws.String(m.Subject),
	}
	body := &ses.Content{
		Charset: aws.String(charset),
		Data:    aws.String(m.Body),
	}

	return &ses.SendEmailInput{
		Destination: &ses.Destination{ToAddresses: []*string{aws.String(m.To)}},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: body,
			},
			Subject: subject,
		},
		Source: aws.String("hello@statistico.io"),
	}
}

func NewSesNotifier(c sesiface.SESAPI) notification.Notifier {
	return &sesNotifier{client: c}
}
