package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"path/filepath"

	"github.com/gradleless/sylote/lib"
	"github.com/gradleless/sylote/lib/pylote"
)

type Data struct {
	Search         []string `json:"search"`
	ID             string   `json:"id"`
	Update         bool     `json:"update"`
	Time           string   `json:"time"`
	DiscordWebhook string   `json:"discordWebhook"`
}

func GetUserInput() (string, string, bool, string, string) {
	var searchStr string
	fmt.Print("Entrez les termes de recherche (séparés par des virgules) : ")
	fmt.Scanln(&searchStr)

	var choice int
	fmt.Print("Voulez-vous ajouter un ID manuellement ou par mail  ? (0/1) : ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Erreur :", err)
		return "", "", false, "", ""
	}

	if choice != 0 && choice != 1 {
		fmt.Println("Erreur : choix doit être 0 ou 1")
		return "", "", false, "", ""
	}

	var id string
	if choice == 0 {

		fmt.Print("Entrez l'ID : ")
		fmt.Scanln(&id)
	}

	if choice == 1 {
		var email string
		fmt.Print("Entrez votre email : ")
		fmt.Scanln(&email)
		pylote.GetCode(email)

		var code string
		fmt.Print("Entrez le code reçu par mail : ")
		fmt.Scanln(&code)
		id = pylote.GetToken(email, code)

		if id == "" {
			fmt.Println("Erreur : impossible de récupérer l'ID")
			return "", "", false, "", ""
		}
	}

	var update bool
	fmt.Print("Mettre à jour ? (true/false) : ")
	_, err = fmt.Scanln(&update)
	if err != nil {
		fmt.Println("Erreur :", err)
		return "", "", false, "", ""
	}

	if reflect.TypeOf(update).Kind() != reflect.Bool {
		fmt.Println("Erreur : update doit être un booléen (true/false)")
		return "", "", false, "", ""
	}

	var timeStr string

	if !update {
		fmt.Print("Entrez l'heure (AAAA-MM-JJ) : ")
		_, err := fmt.Scanln(&timeStr)
		if err != nil {
			fmt.Println("Erreur :", err)
			return "", "", false, "", ""
		}
	}

	var discordWebhook string
	fmt.Print("Entrez l'URL du webhook Discord : ")
	fmt.Scanln(&discordWebhook)

	return searchStr, id, update, timeStr, discordWebhook
}

func CreateData(searchStr string, id string, update bool, timeStr string, discordWebhook string) struct {
	Search         []string `json:"search"`
	ID             string   `json:"id"`
	Update         bool     `json:"update"`
	Time           string   `json:"time"`
	DiscordWebhook string   `json:"discordWebhook"`
} {
	search := strings.Split(searchStr, ",")
	data := struct {
		Search         []string `json:"search"`
		ID             string   `json:"id"`
		Update         bool     `json:"update"`
		Time           string   `json:"time"`
		DiscordWebhook string   `json:"discordWebhook"`
	}{
		Search:         search,
		ID:             id,
		Update:         update,
		Time:           timeStr,
		DiscordWebhook: discordWebhook,
	}
	return data
}

func SaveDataToFile(data struct {
	Search         []string `json:"search"`
	ID             string   `json:"id"`
	Update         bool     `json:"update"`
	Time           string   `json:"time"`
	DiscordWebhook string   `json:"discordWebhook"`
}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	err = os.Mkdir("data", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Create(exPath + "/data/data.json")
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
}

func ReadDataFromFile() Data {
	ex, err := os.Executable()

	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + "/data/data.json")
	if err != nil {
		fmt.Println("Erreur :", err)
		lib.SendNotification("Erreur lors de la lecture du fichier data.json")
		return Data{}
	}
	defer file.Close()

	fileData := make([]byte, 0)
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			break
		}
		fileData = append(fileData, buffer[:n]...)
	}

	data := Data{}
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		fmt.Println("Erreur :", err)
		lib.SendNotification("Erreur lors de la lecture du fichier data.json")
		return Data{}
	}

	return data
}
