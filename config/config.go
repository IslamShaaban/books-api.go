package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var DB_CONFIG DBConfig

// Config represents the top-level structure of your YAML config
type Config struct {
	Production ConfigEnv `yaml:"production"`
	Testing    ConfigEnv `yaml:"testing"`
	Staging    ConfigEnv `yaml:"staging"`
}

// ConfigEnv represents the environment-specific configuration
type ConfigEnv struct {
	Database DBConfig `yaml:"database"`
}

// InitConfig initializes the configuration based on the environment
func InitConfig() {
	// Define a flag to receive the environment argument
	envFlag := "development"

	// Override envFlag if GO_ENV is set
	if goEnv := os.Getenv("GO_ENV"); goEnv != "" {
		envFlag = goEnv
	}

	// Read and parse YAML configuration based on the provided environment
	config, err := loadConfigData(envFlag)
	if err != nil {
		fmt.Printf("Failed to initialize config: %v\n", err)
		return
	}

	// Access the database configuration
	DB_CONFIG = config.Database

	// Initialize the database with the retrieved configuration
	initDB(&DB_CONFIG)
}

// loadConfigData initializes and returns the configuration based on the specified environment
func loadConfigData(env string) (*ConfigEnv, error) {
	// Set up Viper to read the configuration file
	viper.SetConfigFile("config/env.yaml")

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	// Get the environment-specific configuration
	envConfig := viper.Sub(env)
	if envConfig == nil {
		return nil, fmt.Errorf("unknown environment: %s", env)
	}

	// Replace placeholders in the configuration with environment variables
	replaceEnvVars(envConfig)

	// Unmarshal the environment-specific configuration into a ConfigEnv struct
	var configEnv ConfigEnv
	if err := envConfig.Unmarshal(&configEnv); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return &configEnv, nil
}

// replaceEnvVars replaces placeholders in the configuration with environment variables
func replaceEnvVars(envConfig *viper.Viper) {
	for _, k := range envConfig.AllKeys() {
		value := envConfig.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			envVar := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
			envConfig.Set(k, getEnvOrPanic(envVar))
		} else {
			envConfig.Set(k, value)
		}
	}
}

// getEnvOrPanic retrieves an environment variable or panics if not found
func getEnvOrPanic(env string) string {
	res := os.Getenv(env)
	if len(res) == 0 {
		panic("Mandatory env variable not found: " + env)
	}
	return res
}

func LoadEnv() {
	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(fmt.Sprintf("%s/.env", directory))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
