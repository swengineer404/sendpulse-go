package sendpulse

type AutomationService struct {
	Events *AutomationEventService
}

func newAutomationService(client *Client) *AutomationService{
	return &AutomationService{
		Events: newAutomationEventService(client),
	}
}
