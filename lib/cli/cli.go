package cli

import (
	"fmt"
	"os"

	"reflect"
	"strconv"
	"time"

	"github.com/mavenless/sylote/lib"
	"github.com/mavenless/sylote/lib/pylote"
)

func userInputs() {
	searchStr, id, update, timeStr, discordWebhook := GetUserInput()
	data := CreateData(searchStr, id, update, timeStr, discordWebhook)
	SaveDataToFile(data)
}

func Automation() {
	data := ReadDataFromFile()

	var precJobs []pylote.Job
	actualTime := time.Now()
	startUpTime := time.Now()

	pylote.SetAvailability(data.ID, data.Update, data.Time)
	for {
		jobs := pylote.SortJobs(pylote.GetJobs(), data.Search)

		if !reflect.DeepEqual(jobs, precJobs) {
			precJobs = jobs
			lib.SendNotification("Il y a " + strconv.Itoa(len(jobs)) + " nouvelle(s) offre(s) vous correspondant sur Pylote !")
			lib.SendDiscordNotification(data.DiscordWebhook, jobs)
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
		// Add code for installation logic here
		userInputs()
		// TODO
		os.Exit(0)
	case "uninstall":
		// Add code for uninstallation logic here
		// TODO
		os.Exit(0)
	case "update":
		// Add code for update logic here
		// TODO
		userInputs()
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
