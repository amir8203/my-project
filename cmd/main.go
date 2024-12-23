package main

import (
	"log"
	"my-project/src/api"
	"my-project/src/config"
	"my-project/src/data/cache"
	"my-project/src/data/db"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)
}