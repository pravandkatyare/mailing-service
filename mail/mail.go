package mail

// MailAgent interface
type MailAgent interface {
	Send(*Mail) ([]string, error)
}

// Mail struct is mailing Template
type Mail struct {
	From     MailToFrom  `json:"from"`
	Des      Destination `json:"des" binding:"required"`
	Subject  string      `json:"subject" binding:"required"`
	TextBody string      `json:"textBody" binding:"required"`
	HTMLBody string      `json:"HTMLBody"`
	client   string
}

// MailToFrom keeps data of receiver and sender
type MailToFrom struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Destination struct for adding mail recipients
type Destination struct {
	To  []MailToFrom `json:"to" binding:"required"`
	CC  []MailToFrom `json:"cc"`
	BCC []MailToFrom `json:"bcc"`
}

// SendMail function passes the interface to respective Client's Send() function
func (m *Mail) SendMail(client MailAgent) ([]string, error) {
	return client.Send(m)
}
