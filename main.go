package main

import (
	"fmt"

	"github.com/victormoreira92/go-oportunities/config"
	"github.com/victormoreira92/go-oportunities/router"
)

var (
	logger *config.Logger
)


func main()  {
	logger = config.NewLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config initializazion error: %v", err)
		fmt.Print(err)
	}


	router.Initialize()
}