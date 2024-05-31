package main

import (
	"github.com/Maysa87/Gopportunities.git/config"
	"github.com/Maysa87/Gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}
	//initialize router
	router.Initialize()
}
