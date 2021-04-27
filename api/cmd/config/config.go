package config

import "github.com/joho/godotenv"

// Load environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
