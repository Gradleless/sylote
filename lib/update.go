package lib

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/inconshreveable/go-update"
)

func DoUpdate() error {

	switch runtime.GOOS {
	case "windows":

		Update()

	case "darwin":

		Update()

	case "linux":

		Update()

	default:

		fmt.Println("Système d'exploitation non pris en charge.")

	}
	return nil
}

func Update() error {

	resp, err := http.Get("http://localhost:8080/update")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	updateErr := update.Apply(resp.Body, update.Options{})
	if updateErr != nil {
		return updateErr
	}

	return updateErr
}
