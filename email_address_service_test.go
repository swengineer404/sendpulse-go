package sendpulse

import "testing"

func TestEmailAddressService_DeleteEmail(t *testing.T) {
	client := newTestClient()
	if err := client.Email.Address.DeleteEmail("test@go-test.com"); err != nil {
		t.Fatal(err)
	}
}

func TestEmailAddressService_GetEmailLists(t *testing.T) {
	client := newTestClient()
	lists, err := client.Email.Address.GetEmailLists("test@go-test.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(lists)
}
