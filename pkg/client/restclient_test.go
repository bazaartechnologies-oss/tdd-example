package client

import "testing"

func TestGetEmployeeDetails(t *testing.T) {
	var response = make(map[string]interface{})
	url := "https://dummy.restapiexample.com/api/v1/employees"
	h := map[string]string{
		"Authorization": "1234566",
		"Accept":        `application/json`,
	}
	success, response := getEmployees(url, h)

	if success != true {
		t.Errorf("Failed to get response from server: %v", success)

	}
	if response == nil {
		t.Errorf("No response found")
	}

}
