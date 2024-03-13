package lib

import (
	"fmt"
	"log"

	"os"
	"runtime"

	"github.com/jannson/go-autostart"
)

func SetAutoStart() {
	app := &autostart.App{
		Name:        "Sylote",
		DisplayName: "Sylote",
		Exec:        GetExecCommand(),
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
		Exec:        GetExecCommand(),
	}

	if app.IsEnabled() {
		if err := app.Disable(); err != nil {
			log.Fatal(err)
		}
	}
}

func GetExecCommand() []string {
	execCommand := "sylote" // Default executable name

	if runtime.GOOS == "windows" {
		execCommand += ".exe"
	}
	fmt.Println(GetExecPath())
	return []string{GetExecPath()}
}

func GetExecPath() string {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return execPath
}

func CheckEnable() bool {
	app := &autostart.App{
		Name:        "Sylote",
		DisplayName: "Sylote",
		Exec:        GetExecCommand(),
	}
	return app.IsEnabled()
}
