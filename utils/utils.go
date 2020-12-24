package utils

import (
	"github.com/mailjet/mailjet-apiv3-go"
	"github.com/pravandkatyare/mailing-service/logging"
	"github.com/pravandkatyare/mailing-service/mail"
	mjclient "github.com/pravandkatyare/mailing-service/mailingClients/mailjet"
	"github.com/spf13/viper"
)

// GetMailjetTemplate for getting an object of mailjet
func GetMailjetTemplate(mails []mail.Mail) (mjclient.MJClient, error) {
	var mailjetTemplate []mailjet.InfoMessagesV31
	var toRecipient, ccRecipient, bccRecipients *mailjet.RecipientsV31
	for _, m := range mails {
		from, err := getSenderObject()
		if err != nil {
			logging.Errorf("Mailjet registered emails were not set")
			return mjclient.MJClient{}, err
		}
		toRecipient, ccRecipient, bccRecipients = getRecipientObject(m.Des)
		mailjetTemplate = append(mailjetTemplate,
			mailjet.InfoMessagesV31{
				From:     from,
				To:       toRecipient,
				Cc:       ccRecipient,
				Bcc:      bccRecipients,
				Subject:  m.Subject,
				TextPart: m.TextBody,
				HTMLPart: m.HTMLBody,
			})
	}

	return mjclient.MJClient{
		MailTemplate: mailjetTemplate,
	}, nil
}

// function creates mail's sender object
func getSenderObject() (*mailjet.RecipientV31, error) {
	return &mailjet.RecipientV31{
		Email: viper.GetString("mailjet.email"),
		Name:  viper.GetString("mailjet.name"),
	}, nil
}

// function creates the recipient(i.e to, cc, bcc) object
func getRecipientObject(destinationObj mail.Destination) (*mailjet.RecipientsV31, *mailjet.RecipientsV31, *mailjet.RecipientsV31) {

	// for TO recipients
	var toRecipients mailjet.RecipientsV31
	for _, destination := range destinationObj.To {
		toRecipients = append(toRecipients, getRecipient(destination))
	}

	// for CC recipients
	var ccRecipients mailjet.RecipientsV31
	if len(destinationObj.CC) > 0 {
		for _, destination := range destinationObj.CC {
			ccRecipients = append(ccRecipients, getRecipient(destination))
		}
	}

	// for BCC recipients
	var bccRecipients mailjet.RecipientsV31
	if len(destinationObj.BCC) > 0 {
		for _, destination := range destinationObj.BCC {
			bccRecipients = append(bccRecipients, getRecipient(destination))
		}
	}

	return &toRecipients, &ccRecipients, &bccRecipients
}

// function returns an object of independent recipient
func getRecipient(destination mail.MailToFrom) mailjet.RecipientV31 {
	return mailjet.RecipientV31{
		Email: destination.ID,
		Name:  destination.Name,
	}
}
