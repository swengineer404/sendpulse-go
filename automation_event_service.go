package sendpulse

type AutomationEventService struct {
	client *Client
}

func newAutomationEventService(client *Client) *AutomationEventService {
	return &AutomationEventService{
		client: client,
	}
}
