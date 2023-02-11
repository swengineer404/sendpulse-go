package sendpulse

type EmailService struct {
	Address *EmailAddressService
	Mailing *MailingService
}

func newEmailService(client *Client) *EmailService {
	return &EmailService{
		Mailing: newMailingService(client),
		Address: newEmailAddressService(client),
	}
}
