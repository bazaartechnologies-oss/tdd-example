package client

import "testing"

func TestCallWebApp(t *testing.T) {
	var success bool
	success = getEmployees()

	if success != true {
		t.Errorf("Failed to get response from server: %v", success)

	}

}
