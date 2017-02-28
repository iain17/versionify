package main

import (
	logger "github.com/Sirupsen/logrus"
	"os"
	"github.com/Pocketbrain/versionify/examples/http_reverse/api"
)

func main() {
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logger.DebugLevel)

	//Components
	go api.Setup()

	logger.Info("Server is up and running.")

	select {}
}
