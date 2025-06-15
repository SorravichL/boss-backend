package config

import (
  "log"
  "os"
  "github.com/joho/godotenv"
)
type Config struct {
  DBHost, DBPort, DBUser, DBPass, DBName string
}

func Load() *Config {
  //handle error
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file found, relying on environment variables.")
  }
  return &Config{
    DBHost: os.Getenv("DB_HOST"),
    DBPort: os.Getenv("DB_PORT"),
    DBUser: os.Getenv("DB_USER"),
    DBPass: os.Getenv("DB_PASS"),
    DBName: os.Getenv("DB_NAME"),
  }
}