package sgclient

import (
	"fmt"
)

// SGClient struct
type SGClient struct {
}

// Send function has business logic to send mail using SendGrid
func (sg *SGClient) Send() ([]string, error) {
	fmt.Println("SendGrid Client for future enhancement")
	return nil, nil
}
