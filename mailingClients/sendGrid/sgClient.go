package sgclient

import (
	"fmt"

	"github.com/pravandkatyare/mailing-service/mail"
)

// SGClient struct
type SGClient struct {
}

// Send function has business logic to send mail using SendGrid
func (sg *SGClient) Send(mail *mail.Mail) ([]string, error) {
	fmt.Println("SendGrid Client for future enhancement")
	return nil, nil
}
