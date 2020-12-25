package sesclient

import (
	"fmt"
)

// SESClient struct
type SESClient struct {
}

// Send function has business logic to send mail using Amazon SES
func (ses *SESClient) Send() ([]string, error) {
	fmt.Println("Amazon SES Client for future enhancement")
	return nil, nil
}
