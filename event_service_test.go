package sendpulse

import "testing"

func TestEventService_Invoke(t *testing.T) {
	client := newTestClient()
	if err := client.Events.Invoke("kill_no_purchase", "test@go-test.com"); err != nil {
		t.Fatal(err)
	}
}
