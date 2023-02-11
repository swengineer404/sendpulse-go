package sendpulse

import (
	"fmt"
	"net/url"
)

type EmailAddressService struct {
	client *Client
}

type EmailList struct {
	ID      int    `json:"list_id"`
	Name    string `json:"list_name"`
	AddDate string `json:"add_date"`
	Source  string `json:"source"`
}

func newEmailAddressService(client *Client) *EmailAddressService {
	return &EmailAddressService{
		client: client,
	}
}

func (s *EmailAddressService) GetEmailLists(email string) (result []EmailList, err error) {
	_, err = s.client.Send("GET", fmt.Sprintf("/emails/%s/details", email), nil, &result, true)

	return result, err
}

func (s *EmailAddressService) DeleteEmail(email string) error {
	_, err := s.client.Send("DELETE", fmt.Sprintf("/emails/%s", url.PathEscape(email)), nil, nil, true)

	return err
}
