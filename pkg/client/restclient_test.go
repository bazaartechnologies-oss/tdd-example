package client

import "testing"

func TestGetEmployeeDetails(t *testing.T) {
	var response = make(map[string]interface{})
	success, response := getEmployees()

	if success != true {
		t.Errorf("Failed to get response from server: %v", success)

	}
	if response == nil {
		t.Errorf("No response found")
	}

}
