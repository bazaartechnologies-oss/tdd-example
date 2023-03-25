package client

import (
	"log"
	"net/http"
)

func getEmployees() (success bool) {

	_, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		log.Printf("Unable to call the server: %v", err)
		return success
	}
	success = true
	return success
}
