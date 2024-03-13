package pylote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Origin", "moz-extension://18443650-c9cb-407e-83ea-b7e0f5465571")
	req.Header.Add("DNT", "1")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "cross-site")
	req.Header.Add("TE", "trailers")

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

	var body map[string]interface{}
	err = json.Unmarshal(bodyBytes.Bytes(), &body)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	if body["msg"] == "OK" {
		fmt.Println("Availability set")
	} else {
		fmt.Println("Error while setting availability")
	}
}
