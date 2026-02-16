package main

import (
	"github.com/MatheusMikio/eduGuard_api/config"
	"github.com/MatheusMikio/eduGuard_api/internal/routes"
)

func main() {
	logger := config.GetLogger("main")

	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	db := config.GetDb()

	routes.Init(db)
}
