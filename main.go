package main

import (
	"github.com/edvaldocebol/projeto-Golang/config"
	"github.com/edvaldocebol/projeto-Golang/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil { 
		logger.Errorf("config initialization error %v", err)
		return
	}
	router.Inisialize() 
}
