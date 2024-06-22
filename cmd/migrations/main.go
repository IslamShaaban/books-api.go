package main

import (
	"books-api/config"
	_ "books-api/db/migrations"
	"context"
	"flag"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	ctx   context.Context
)

func main() {

	// Checking Arguments
	if len(os.Args) == 1 {
		log.Fatalf("No commands sent to migrate")
	}

	// Parsing All Arguments Flags
	flags.Parse(os.Args[1:])
	args := flags.Args()

	// Load Config Data
	config.LoadEnv()
	config.InitConfig()

	// Set Command Running
	command := "status"
	if len(args) > 0 {
		command = args[0]
	}

	// Set Database Dialect
	if err := goose.SetDialect(config.DB_CONFIG.Adapter); err != nil {
		log.Fatalf("goose error setting dialect: %v", err)
	}

	// Run Migration Command
	runMigrationFile(command, args)
}

func runMigrationFile(command string, args []string) {
	directory, err := os.Getwd()

	// Create a context
	ctx := context.Background()

	if err != nil {
		panic(err)
	}
	migrationDirectory := fmt.Sprintf("%s/db/migrations", directory)

	log.Println("Running Migration Script in Progress...")
	if command != "" {
		log.Printf("Running: %s with values %v", command, args[1:])
		if err := goose.RunContext(ctx, command, config.SQL_DB, migrationDirectory, args[1:]...); err != nil {
			log.Fatalf("goose run error: %v", err)
		}
		return
	}
}
