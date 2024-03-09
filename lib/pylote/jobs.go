package pylote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Job struct {
	URL          string `json:"URL,omitempty"`
	TJM          string `json:"TJM,omitempty"`
	City         string `json:"city,omitempty"`
	Date         string `json:"Date,omitempty"`
	JoursSemaine string `json:"Jours/semaine,omitempty"`
	Plateforme   string `json:"Plateforme,omitempty"`
	Remote       string `json:"remote,omitempty"`
	Durée        string `json:"Durée,omitempty"`
	DuréeMois    string `json:"Durée_mois,omitempty"`
	Title        string `json:"Title,omitempty"`
	ID           string `json:"id,omitempty"`
}

func GetJobs() []Job {

	url := "https://api-p.pylote.io/jobs/"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	var bodyBytes bytes.Buffer
	_, err = io.Copy(&bodyBytes, res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var jobs []Job
	err = json.Unmarshal(bodyBytes.Bytes(), &jobs)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return nil
	}

	return jobs
}

func SortJobs(jobs []Job, words []string) []Job {
	var sortedJobs []Job
	for _, job := range jobs {
		var found bool
		for _, word := range words {
			for _, value := range []string{job.TJM, job.City, job.Date, job.JoursSemaine, job.Plateforme, job.Remote, job.Durée, job.Title, job.DuréeMois} {
				if strings.Contains(value, word) {
					found = true
					break
				} else {
					found = false
				}
			}
		}
		if found {
			sortedJobs = append(sortedJobs, job)
		}
	}
	return sortedJobs
}
