package mjclient

import (
	mailjet "github.com/mailjet/mailjet-apiv3-go"
	"github.com/pravandkatyare/mailing-service/logging"
)

const (
	successResponse = "success"
)

// MJClient struct
type MJClient struct {
	MailTemplate []mailjet.InfoMessagesV31
}

// Send function has business logic to send mail using MailJet
func (m MJClient) Send() ([]string, error) {
	logging.Infof("In MJClient: Send(), sending mail through Mail Jet")

	mailjetClient := mailjet.NewMailjetClient(getPublicKey(), getPrivateKey())
	messages := mailjet.MessagesV31{Info: m.MailTemplate}

	res, err := mailjetClient.SendMailV31(&messages)
	if err != nil {
		logging.Errorf("Error sending mail: %s", err)
		return nil, err
	}

	var mailFailures []string
	for _, res := range res.ResultsV31 {
		if res.Status != successResponse {
			for _, to := range res.To {
				mailFailures = append(mailFailures, to.Email)
			}
		}

	}

	// no mail failures means all the mails are sent successfully
	if len(mailFailures) == 0 {
		logging.Infof("Mail sent successfully to recipient/s", res)
		return nil, nil
	} else {
		logging.Infof("Mail not sent to recipient/s", res)
		return mailFailures, nil
	}
}
