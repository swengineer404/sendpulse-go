package sendpulse

import "fmt"

type MailingService struct {
	client *Client
}

func newMailingService(client *Client) *MailingService {
	return &MailingService{
		client: client,
	}
}

func (s *MailingService) SingleOptIn(listID int, emails ...string) error {
	_, err := s.client.Send("POST", fmt.Sprintf("/addressbooks/%d/emails", listID), emails, nil, true)

	return err
}
