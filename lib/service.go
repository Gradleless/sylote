package lib

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kardianos/service"
)

type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	close(p.exit)
	return nil
}

func (p *program) run() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("Programme en cours d'exécution...")
		case <-p.exit:
			return
		}
	}
}

func setService() {
	prg := &program{exit: make(chan struct{})}

	svcConfig := &service.Config{
		Name:        "Sylote",
		DisplayName: "Sylote Service",
		Description: "Service permettant de trouver des boulots vous intéressant et de mettre à jour votre statut automatiquement",
	}

	s, err := service.New(prg, svcConfig)
	if err != nil {
		fmt.Println("Erreur lors de la création du service:", err)
		os.Exit(1)
	}

	// Gestion des signaux pour permettre une fermeture propre
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		fmt.Printf("Signal reçu: %v\n", sig)
		s.Stop()
	}()

	// Démarrage du service
	if err := s.Run(); err != nil {
		fmt.Println("Erreur lors de l'exécution du service:", err)
	}
}

func deleteService() {
	svcConfig := &service.Config{
		Name: "Sylote",
	}

	s, err := service.New(nil, svcConfig)
	if err != nil {
		fmt.Println("Erreur lors de la création du service:", err)
		os.Exit(1)
	}

	if err := s.Uninstall(); err != nil {
		fmt.Println("Erreur lors de la suppression du service:", err)
	} else {
		fmt.Println("Service supprimé avec succès.")
		SendNotification("Service supprimé avec succès.")
	}
}
