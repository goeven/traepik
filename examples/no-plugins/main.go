package main

import (
	"log"

	"github.com/goeven/traepik/pkg/cmd"
)

func main() {
	traefikCmd, err := cmd.New(nil)
	if err != nil {
		log.Fatalf("Failed creating command traefik: %v", err)
	}

	if err := traefikCmd.Execute(); err != nil {
		log.Fatalf("traefik: %v", err)
	}
}
