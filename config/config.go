package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config - struct of config. variables
type Config struct {
	DatabaseHost           string
	DatabasePort           string
	DatabaseName           string
	DatabasePassword       string
	DatabaseUsername       string
	MagacoinPaymentURL     string
	PrivateKey             string
	TargetAccount          int
	DatabaseMaxIdle        int
	DatabaseIdleTimeout    int
	MagacoinPaymentTimeout int
	DatabaseMaxCon         int
}

var (
	// Env - export environment variables
	Env = GetConfig()
)

// LoadEnvs - load environment variables
func LoadEnvs() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Running application without .env file")
	}
}

// GetConfig - get values of environment variables
func GetConfig() Config {
	var config Config
	LoadEnvs()
	config.DatabaseHost = os.Getenv("DATABASE_HOST")
	config.DatabasePort = os.Getenv("DATABASE_PORT")
	config.DatabaseName = os.Getenv("DATABASE_NAME")
	config.DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	config.DatabaseUsername = os.Getenv("DATABASE_USERNAME")
	config.DatabaseMaxIdle = GetIntFromEnv("DATABASE_MAX_IDLE")
	config.DatabaseIdleTimeout = GetIntFromEnv("DATABASE_IDLE_TIMEOUT")
	config.DatabaseMaxCon = GetIntFromEnv("DATABASE_MAX_CON")
	config.MagacoinPaymentURL = os.Getenv("MAGACOIN_PAYMENT_URL")
	config.MagacoinPaymentTimeout = GetIntFromEnv("MAGACOIN_PAYMENT_TIMEOUT")
	config.PrivateKey = os.Getenv("PRIVATE_KEY")
	config.TargetAccount = GetIntFromEnv("TARGET_ACCOUNT")
	return config
}

// GetIntFromEnv - parse int variables values
func GetIntFromEnv(key string) int {
	env := os.Getenv(key)
	intEnv, err := strconv.Atoi(env)
	if err != nil {
		return 0
	}
	return intEnv
}

// GetBoolFromEnv - parse boolean variables values
func GetBoolFromEnv(key string) bool {
	env := os.Getenv(key)
	boolEnv, err := strconv.ParseBool(env)
	if err != nil {
		return false
	}
	return bool(boolEnv)
}
