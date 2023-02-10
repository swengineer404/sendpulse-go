package sendpulse

type EmailAddressService struct {
	client *Client
}

func newEmailAddressService(client *Client) *EmailAddressService {
	return &EmailAddressService{
		client: client,
	}
}
