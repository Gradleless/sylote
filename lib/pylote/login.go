package pylote

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Token struct {
	ID    string `json:"id"`
	Check bool   `json:"check"`
}

func Login() (string, string) {
	fmt.Print("Email: ")
	var email string
	fmt.Scanln(&email)
	fmt.Print("Code d'accès (dans votre boite mail): ")
	var code string
	fmt.Scanln(&code)
	return email, code
}

func GetCode(email string) {

	url := "https://api-p.pylote.io/freelance/set_login_code"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"mail": "%s","resend":false}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Origin", "moz-extension://422b98c5-8fbb-4b4c-801c-e2783df5e9b6")
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
}

func GetToken(email string, code string) string {
	url := "https://api-p.pylote.io/freelance/check_code"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"mail": "%s", "code": "%s"}`, email, code))
	fmt.Println(payload)
	fmt.Println("gnan")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:123.0) Gecko/20100101 Firefox/123.0")
	req.Header.Add("Accept", "application/json, */*")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Origin", "moz-extension://422b98c5-8fbb-4b4c-801c-e2783df5e9b6")
	req.Header.Add("Authorization", "Bearer 74381179-6e82-44c5-a0a5-c72d9181149a")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	var bodyBytes bytes.Buffer
	_, err = io.Copy(&bodyBytes, res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(bodyBytes.String())

	return ""
}
