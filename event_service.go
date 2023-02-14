package sendpulse

import "fmt"

type EventService struct {
	client *Client
}

func newEventService(client *Client) *EventService {
	return &EventService{
		client: client,
	}
}

type InvokeEventParams struct {
	Email string `json:"email"`
}

func (s *EventService) Invoke(eventName string, email string) error {
	params := &InvokeEventParams{
		Email: email,
	}

	_, err := s.client.InvokeEvent("POST", fmt.Sprintf("/events/name/%s", eventName), params, nil)

	return err
}
