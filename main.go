package main

import (
	"github.com/stgonzles/drillit-api/config"
	"github.com/stgonzles/drillit-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Failed to load db config: %v", err)
		return
	}

	router.Initialize()
}