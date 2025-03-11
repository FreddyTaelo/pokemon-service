// config/config.go
package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PokeAPIURL         string
	Port               string
	Address            string
	RateLimitPerSecond string
	RateLimitBurst     string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables.")
	}

	return &Config{
		PokeAPIURL:         getEnv("POKEAPI_URL", "https://pokeapi.co/api/v2"),
		Port:               getEnv("PORT", "5000"),
		Address:            getEnv("ADDR", "0.0.0.0"),
		RateLimitPerSecond: getEnv("RATE_PER_SECOND", "20"),
		RateLimitBurst:     getEnv("RATE_BURST", "100"),
	}
} /*LoadConfig()*/

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
} /*getEnv()*/
