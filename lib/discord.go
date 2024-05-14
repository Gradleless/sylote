package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/gradleless/sylote/lib/pylote"
)

type DiscordWebhookPayload struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       int     `json:"color"`
	Fields      []Field `json:"fields"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func SendDiscordNotification(webhookURL string, jobs []pylote.Job) {

	urls := []Field{}
	embeds := []Embed{}
	for _, job := range jobs {

		urls = append(urls, Field{
			Name:  job.Title,
			Value: job.URL,
		})

		for i, value := range []string{job.TJM, job.City, job.Date, job.JoursSemaine, job.Plateforme, job.Remote, job.Durée, job.DuréeMois} {
			if value == "" {
				value = "N/A"
			}

			if reflect.TypeOf(job).Field(i+1).Name == "Date" {
				value = value[:10]
			}

			urls = append(urls, Field{
				Name:  reflect.TypeOf(job).Field(i + 1).Name,
				Value: value,
			})
		}

		if len(urls) >= 16 {
			embeds = append(embeds, Embed{
				Title:  "Sylote",
				Color:  5814783,
				Fields: urls,
			})
			urls = []Field{}
		}
	}

	payload := DiscordWebhookPayload{
		Embeds: embeds,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
}
