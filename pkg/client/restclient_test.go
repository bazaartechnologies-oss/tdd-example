package client

import "testing"

func TestGetEmployeeDetails(t *testing.T) {
	var response = make(map[string]interface{})
	url := "https://dummy.restapiexample.com/api/v1/employees"
	success, response := getEmployees(url)

	if success != true {
		t.Errorf("Failed to get response from server: %v", success)

	}
	if response == nil {
		t.Errorf("No response found")
	}

}
