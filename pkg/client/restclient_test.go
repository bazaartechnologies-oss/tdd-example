package client

import (
	"net/http"
	"testing"
)

func TestGetEmployeeDetails(t *testing.T) {
	httpRequest := HTTPRequest{
		URL: "https://dummy.restapiexample.com/api/v1/employees",
		Headers: map[string]string{
			"Authorization": "asd",
			"Accept":        `application/json`,
		},
		Method: http.MethodGet,
	}
	success, response := httpRequest.Client()

	if success != true {
		t.Errorf("Failed to get response from server: %v", success)

	}
	if response == nil {
		t.Errorf("No response found")
	}

}

func TestCreateEmployee(t *testing.T) {
	httpRequest := HTTPRequest{
		Headers: map[string]string{
			"Content-Type": `application/json`,
		},
		URL:    "https://dummy.restapiexample.com/api/v1/create",
		Method: http.MethodPost,
		Body:   []byte(`{"name":"test","salary":"123","age":"23"}`),
	}

	_, respone := httpRequest.Client()

	if respone["status"] != "success" {
		t.Errorf("Failed to create employee")
	}

}
