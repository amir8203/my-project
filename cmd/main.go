package main

import (
	"log"
	"my-project/src/config"
	"my-project/src/data/cache"
	"my-project/src/data/db"
)

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
}