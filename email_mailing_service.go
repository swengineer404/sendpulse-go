package sendpulse

import (
	"fmt"
)

type MailingService struct {
	client *Client
}

func newMailingService(client *Client) *MailingService {
	return &MailingService{
		client: client,
	}
}

type MailingList struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *MailingService) GetLists(limit, offset int) (result []MailingList, err error) {
	_, err = s.client.Send("GET", fmt.Sprintf("/addressbooks?limit=%d&offset=%d", limit, offset), nil, &result, true)

	return result, err
}

func (s *MailingService) SingleOptIn(listID int, emails ...string) error {
	body := map[string][]string{
		"emails": emails,
	}

	_, err := s.client.Send("POST", fmt.Sprintf("/addressbooks/%d/emails", listID), body, nil, true)

	return err
}

func (s *MailingService) DeleteEmails(listID int, emails ...string) error {
	body := map[string][]string{
		"emails": emails,
	}

	_, err := s.client.Send("DELETE", fmt.Sprintf("/addressbooks/%d/emails", listID), body, nil, true)

	return err
}
