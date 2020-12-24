package sesclient

import (
	"fmt"

	"github.com/pravandkatyare/mailing-service/mail"
)

// SESClient struct
type SESClient struct {
}

// Send function has business logic to send mail using Amazon SES
func (ses *SESClient) Send(mail *mail.Mail) ([]string, error) {
	fmt.Println("SESClient")
	return nil, nil
}
