package lib

import (
	"github.com/gen2brain/beeep"
)

func SendNotification(message string) {
	err := beeep.Notify("Sylote", message, "")
	if err != nil {
		panic(err)
	}
}
