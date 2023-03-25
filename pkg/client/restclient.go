package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getEmployees() (success bool, response map[string]interface{}) {

	r, err := http.Get("https://dummy.restapiexample.com/api/v1/employees")
	if err != nil {
		log.Printf("Unable to call the server: %v", err)
		return success, response
	}
	success = true
	if err != nil {
		return false, nil
	}
	rb, err := io.ReadAll(r.Body)

	if err != nil {
		log.Printf("Unable to convert into byte: %v", err)
	}
	err = json.Unmarshal(rb, &response)
	if err != nil {
		log.Printf("failed conversion: %v", err)
	}
	fmt.Println(response)
	return success, response
}
