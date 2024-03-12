package util

import "github.com/joho/godotenv"

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		GetGrpcLoggerV2().Fatalf("Error loading .env file: %v", err.Error())
	}
}
