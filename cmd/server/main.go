package main

import (
	"books-api/config"
	"books-api/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	// Load All Env Variables
	config.LoadEnv()

	// Initialize Config and
	config.InitConfig()

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRouter(r, config.GetDBConnection())

	// Run Server
	serve(r)
}

func serve(r *gin.Engine) {
	var (
		ok       bool
		hostname string
	)
	hostname, ok = os.LookupEnv("SERVER_HOSTNAME")
	if !ok {
		hostname = "localhost"
	}
	connectionString := fmt.Sprintf("%s:%s", hostname, os.Getenv("SERVER_PORT"))
	r.Run(connectionString)
}
