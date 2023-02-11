package sendpulse

import "testing"

func TestMailingService_GetLists(t *testing.T) {
	client := newTestClient()
	lists, err := client.Email.Mailing.GetLists(50, 0)
	if err != nil {
		t.Fatal(err)
	}

	for _, list := range lists {
		t.Log(list.Name)
	}
}

func TestMailingService_SingleOptIn(t *testing.T) {
	client := newTestClient()
	if err := client.Email.Mailing.SingleOptIn(79104, "test@go-test.com"); err != nil {
		t.Fatal(err)
	}
}

func TestMailingService_DeleteEmails(t *testing.T) {
	client := newTestClient()
	if err := client.Email.Mailing.DeleteEmails(79104, "test@go-test.com"); err != nil {
		t.Fatal(err)
	}
}
