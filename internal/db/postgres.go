package db

import (
  "fmt"
  "log"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "boss-backend/internal/config"
)

func NewGormDB(cfg *config.Config) *gorm.DB {
  dsn := fmt.Sprintf(
    "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
    cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
  )
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("failed to connect database: %v", err)
  }
  return db
}
