package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getEmployees(url string, headers map[string]string) (success bool, response map[string]interface{}) {

	c := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("error %v", err)
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	r, err := c.Do(req)

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
