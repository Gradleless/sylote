package cli

import (
	"fmt"
	"os"

	"strconv"
	"time"

	"github.com/mavenless/sylote/lib"
	"github.com/mavenless/sylote/lib/pylote"
)

func ArrayElementEquals(array1 []pylote.Job, array2 []pylote.Job) bool {
	if len(array1) != len(array2) {
		return false
	}
	for i := range array1 {
		if array1[i] != array2[i] {
			return false
		}
	}
	return true
}

func userInputs() {
	searchStr, id, update, timeStr, discordWebhook := GetUserInput()
	data := CreateData(searchStr, id, update, timeStr, discordWebhook)
	SaveDataToFile(data)
}

func Automation() {
	lib.SendNotification("Démarrage de Sylote")
	data := ReadDataFromFile()

	var precJobs []pylote.Job
	actualTime := time.Now()
	startUpTime := time.Now()

	pylote.SetAvailability(data.ID, data.Update, data.Time)
	for {
		jobs := pylote.SortJobs(pylote.GetJobs(), data.Search)

		if !ArrayElementEquals(jobs, precJobs) {
			lib.SendNotification("Il y a " + strconv.Itoa(len(jobs)-len(precJobs)) + " nouvelle(s) offre(s) vous correspondant sur Pylote !")
			lib.SendDiscordNotification(data.DiscordWebhook, jobs)
			precJobs = jobs
		}

		if time.Since(startUpTime).Hours() >= 24 && data.Update {
			pylote.SetAvailability(data.ID, data.Update, data.Time)
			startUpTime = time.Now()
		}

		if time.Since(actualTime).Hours() >= 1 {
			actualTime = time.Now()
		} else {
			time.Sleep(1 * time.Hour)
		}
	}
}

func Cli(args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide a valid command: install, uninstall, or update")
		return
	}

	command := args[1]
	switch command {
	case "install":
		lib.SetAutoStart()
		userInputs()
		os.Exit(0)
	case "uninstall":
		fmt.Println("Désinstallation de l'application")
		lib.UnsetAutoStart()
		os.Exit(0)
	case "update":
		userInputs()
		os.Exit(0)
	case "help":
		fmt.Println("Usage: sylote [command]")
		fmt.Println("Commands:")
		fmt.Println("  install   Installer l'application")
		fmt.Println("  uninstall Déinstaller l'application")
		fmt.Println("  update    Mettre à jour vos données")
		os.Exit(0)

	case "start":
		Automation()
	default:
		fmt.Println("Commande non reconnue. Utilisez 'sylote help' pour obtenir de l'aide.")
	}
}
