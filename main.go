package main

import (
	"strconv"

	"github.com/mavenless/sylote/lib"
	"github.com/mavenless/sylote/lib/pylote"
)

func main() {

	jobs := pylote.SortJobs(pylote.GetJobs(), []string{"Développeur", "fullstack"})
	lib.SendNotification("Il y a " + strconv.Itoa(len(jobs)) + " nouvelle(s) offre(s) vous correspondant sur Pylote !")
	pylote.SetAvailability("id pylote", true, "")
	lib.SendDiscordNotification("https://discord.com/api/webhooks/1216095231968018513/ZF5T1IiE53eNExYg8iiBgljkEJ9dyOnZQHmwhMlrmUsgk1YaSkmhHRI_3_6-NlBS2LE8", "Il y a "+strconv.Itoa(len(jobs))+" nouvelle(s) offre(s) vous correspondant sur Pylote !", jobs)
}
