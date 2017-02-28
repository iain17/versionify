package main

import (
	"github.com/Pocketbrain/versionify/examples/http/api"
	logger "github.com/Sirupsen/logrus"
	"os"
)

func main() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.DebugLevel)

	//Components
	go api.Setup()

	logger.Info("Server is up and running.")

	select {}
}
