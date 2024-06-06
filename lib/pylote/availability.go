package pylote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func SetAvailability(id string, available bool, date string) {
	if available {
		date = ""
	}

	url := "https://api-p.pylote.io/availability/" + id
	method := "PUT"

	payload := strings.NewReader(fmt.Sprintf("{\"available\": %t, \"availabilityDate\": \"%s\"}", available, date))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("User-Agent", "Sylote")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	var bodyBytes bytes.Buffer
	_, err = io.Copy(&bodyBytes, res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var responseBody map[string]interface{}
	err = json.Unmarshal(bodyBytes.Bytes(), &responseBody)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSONe:", err)
		return
	}

	if responseBody["msg"] == "OK" {
		fmt.Println("Availability set")
	} else {
		fmt.Println("Error while setting availability")
	}

	UpdateAvailability(id)
}

func UpdateAvailability(id string) {
	url := "https://api-p.pylote.io/logging/update"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("{\"id\": \"%s\", \"variable\": \"m\", \"value\": \"%s\"}", id, time.Now().Format(time.RFC3339Nano)))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("User-Agent", "Sylote")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Connection", "keep-alive")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	var bodyBytes bytes.Buffer
	_, err = io.Copy(&bodyBytes, res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bodyBytes.String())
	var body map[string]interface{}
	err = json.Unmarshal(bodyBytes.Bytes(), &body)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSONfjfjfh:", err)
		return
	}

	if body["msg"] == "OK" {
		fmt.Println("Logging updated")
	} else {
		fmt.Println("Error while updating availability")
	}
}
