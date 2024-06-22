package seeder

import (
	"books-api/config"
	// "books-api/db/seeds"
	"time"
)

func generateSeed() {
	loc, err := time.LoadLocation("UTC")
	if err == nil {
		time.Local = loc
	}
	config.LoadEnv()
	config.InitConfig()
	// seeds.RunAllSeeders()
}
