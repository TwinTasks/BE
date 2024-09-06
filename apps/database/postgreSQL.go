package database

import (
	"fmt"
	"log"
	"twintask/apps/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Setting) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port);

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{});

	if err != nil {
		log.Fatal("Error config Database", err.Error());
	};

	return DB;
}