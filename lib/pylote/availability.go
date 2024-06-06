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
	UpdateAvailability(id)

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
		fmt.Println("Erreur lors du décodage JSON (SetAvailability) :", err)
		return
	}

	if responseBody["msg"] == "OK" {
		fmt.Println("Availability set")
	} else {
		fmt.Println("Error while setting availability")
	}
}

func UpdateAvailability(id string) {
	url := "https://api-p.pylote.io/logging/update"
	method := "POST"

	logs, err := GetLogs(id)
	if err != nil {
		fmt.Println("Error fetching logs:", err)
		return
	}

	variablesToUpdate := map[string]interface{}{
<<<<<<< HEAD
		"dd": logs.Dd + 1,
		"cc": logs.Cc + 1,
		"a":  logs.A + 1,
		"m":  time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
	}

	for key, value := range variablesToUpdate {
		payload := strings.NewReader(fmt.Sprintf("{\"id\": \"%s\", \"variable\": \"%s\", \"value\": \"%v\"}", logs.Id, key, value))
=======
		"a":  logs.A + 1,
		"m":  time.Now().Format(time.RFC3339Nano),
		"cc": logs.Cc + 1,
		"dd": logs.Dd + 1,
	}

	for key, value := range variablesToUpdate {
		payload := strings.NewReader(fmt.Sprintf("{\"id\": \"%s\", \"variable\": \"%s\", \"value\": \"%v\"}", id, key, value))
>>>>>>> d6eec15 (pylote change 4 variables, why ???)

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

		var body map[string]interface{}
		err = json.Unmarshal(bodyBytes.Bytes(), &body)
		if err != nil {
			fmt.Println("Erreur lors du décodage JSON (UpdateAvailability) :", err)
			return
		}

		if body["msg"] == "OK" {
<<<<<<< HEAD
			fmt.Printf("Logging updated for variable %s\n %v", key, value)
=======
			fmt.Printf("Logging updated for variable %s\n", key)
>>>>>>> d6eec15 (pylote change 4 variables, why ???)
		} else {
			fmt.Printf("Error while updating variable %s\n", key)
		}
	}
}

func GetLogs(id string) (LogData, error) {
<<<<<<< HEAD
<<<<<<< HEAD

=======
	// ***REMOVED***4
>>>>>>> d6eec15 (pylote change 4 variables, why ???)
=======

>>>>>>> 95aa76e (io)
	url := "https://api-p.pylote.io/logging/getLogs/" + id
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return LogData{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return LogData{}, err
	}
	defer res.Body.Close()

	var bodyBytes bytes.Buffer
	_, err = io.Copy(&bodyBytes, res.Body)
	if err != nil {
		fmt.Println(err)
		return LogData{}, err
	}

	var logData LogData
	err = json.Unmarshal(bodyBytes.Bytes(), &logData)
	if err != nil {
		fmt.Println("Error decoding JSON (GetLogs):", err)
		return LogData{}, err
	}

<<<<<<< HEAD
=======
	fmt.Println("Log Data:", logData)
>>>>>>> d6eec15 (pylote change 4 variables, why ???)
	return logData, nil
}

// Made it with extra cause it'll be used in the future, maybe
type LogData struct {
	A                  int           `json:"a"`
	B                  int           `json:"b"`
	C                  int           `json:"c"`
	M                  string        `json:"m"`
	D                  int           `json:"d"`
	E                  int           `json:"e"`
	F                  int           `json:"f"`
	G                  int           `json:"g"`
	H                  int           `json:"h"`
	N                  string        `json:"n"`
	I                  int           `json:"i"`
	J                  int           `json:"j"`
	K                  string        `json:"k"`
	T                  string        `json:"t"`
	L                  int           `json:"l"`
	Z                  string        `json:"z"`
	P                  int           `json:"p"`
	U                  string        `json:"u"`
	Cc                 int           `json:"cc"`
	Dd                 int           `json:"dd"`
	Gg                 bool          `json:"gg"`
	Id                 string        `json:"id"`
	Hh                 []string      `json:"hh"`
	ExcludeFromUsers   []interface{} `json:"Exclude (from Users)"`
	W                  string        `json:"w"`
	Y                  string        `json:"y"`
	OpenToCDIFromUsers []interface{} `json:"Open_to_CDI (from Users)"`
	PrenomNomFromUsers []string      `json:"Prénom Nom (from Users)"`
	PhoneFromUsers     []string      `json:"phone (from Users)"`
	MailPylote         []string      `json:"Mail Pylote"`
	X                  string        `json:"x"`
	DiffOuvDispo       int           `json:"Diff ouv dispo"`
	Oo                 int           `json:"oo"`
	Pp                 int           `json:"pp"`
	Qq                 bool          `json:"qq"`
	O                  []interface{} `json:"o"`
}
