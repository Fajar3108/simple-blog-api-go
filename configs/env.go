package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type envStruct struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

var Env *envStruct = nil

func init() {
	loadEnvFile()
	checkRequiredVars()

	Env = &envStruct{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
	}
}

func loadEnvFile() {
	if err := godotenv.Load(); err != nil {
		cwd, _ := os.Getwd()
		fmt.Println("CWD: ", cwd)
		log.Fatalf("Error loading .env file or missing environment variable: %v", err)
	}
}

func checkRequiredVars() {
	requiredEnvVars := []string{"DB_USER", "DB_HOST", "DB_PORT", "DB_NAME"}

	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatal("missing environment variable:", envVar)
		}
	}
}
