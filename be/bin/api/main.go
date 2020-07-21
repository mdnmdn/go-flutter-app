package main

import (
	"backend"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("wow")
	backend.StartServer()
}
