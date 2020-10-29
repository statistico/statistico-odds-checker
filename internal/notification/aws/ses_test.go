package aws_test

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/statistico/statistico-odds-checker/internal/notification"
	taws "github.com/statistico/statistico-odds-checker/internal/notification/aws"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestSesNotifier_SendMessage(t *testing.T) {
	t.Run("successfully sends email via AWS SES client and returns success result", func(t *testing.T) {
		t.Helper()

		client := new(taws.MockSesClient)
		notifier := taws.NewSesNotifier(client)

		output := ses.SendEmailOutput{MessageId: aws.String("12345")}

		email := mock.MatchedBy(func(s *ses.SendEmailInput) bool {
			assert.Equal(t, []*string{aws.String("joe@hello.com")}, s.Destination.ToAddresses)
			assert.Equal(t, "hello@statistico.io", *s.Source)
			assert.Equal(t, "UTF-8", *s.Message.Subject.Charset)
			assert.Equal(t, "Hello from Statistico", *s.Message.Subject.Data)
			assert.Equal(t, "UTF-8", *s.Message.Body.Html.Charset)
			assert.Equal(t, "Message Body", *s.Message.Body.Html.Data)
			return true
		})

		client.On("SendEmail", email).Return(&output, nil)

		m := notification.Message{
			To:      "joe@hello.com",
			From:    "hello@statistico.io",
			Subject: "Hello from Statistico",
			Body:    "Message Body",
		}

		result := notifier.SendMessage(&m)

		assert.Equal(t, "SUCCESS", result.Status)
		assert.Equal(t, "12345", result.Message)
	})

	t.Run("fails to sends email via AWS SES client and returns fail result", func(t *testing.T) {
		t.Helper()

		client := new(taws.MockSesClient)
		notifier := taws.NewSesNotifier(client)

		email := mock.MatchedBy(func(s *ses.SendEmailInput) bool {
			assert.Equal(t, []*string{aws.String("joe@hello.com")}, s.Destination.ToAddresses)
			assert.Equal(t, "hello@statistico.io", *s.Source)
			assert.Equal(t, "UTF-8", *s.Message.Subject.Charset)
			assert.Equal(t, "Hello from Statistico", *s.Message.Subject.Data)
			assert.Equal(t, "UTF-8", *s.Message.Body.Html.Charset)
			assert.Equal(t, "Message Body", *s.Message.Body.Html.Data)
			return true
		})

		client.On("SendEmail", email).Return(&ses.SendEmailOutput{}, errors.New("cannot send"))

		m := notification.Message{
			To:      "joe@hello.com",
			From:    "hello@statistico.io",
			Subject: "Hello from Statistico",
			Body:    "Message Body",
		}

		result := notifier.SendMessage(&m)

		assert.Equal(t, "FAIL", result.Status)
		assert.Equal(t, "cannot send", result.Message)
	})
}
