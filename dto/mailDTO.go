package dto

import "github.com/pravandkatyare/mailing-service/mail"

// MailDTO for getting
type MailDTO struct {
	Client string      `json:"client" binding:"required"`
	Mail   []mail.Mail `json:"mail" binding:"required"`
}
