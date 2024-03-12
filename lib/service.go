package lib

import (
	"log"

	"github.com/emersion/go-autostart"
)

func SetAutoStart() {
	app := &autostart.App{
		Name:        "Sylote",
		DisplayName: "Sylote",
		Exec:        []string{"sylote"},
	}
	if !app.IsEnabled() {
		if err := app.Enable(); err != nil {
			log.Fatal(err)
		}
	}
}

func UnsetAutoStart() {
	app := &autostart.App{
		Name:        "Sylote",
		DisplayName: "Sylote",
		Exec:        []string{"sylote"},
	}
	if app.IsEnabled() {
		if err := app.Disable(); err != nil {
			log.Fatal(err)
		}
	}
}
