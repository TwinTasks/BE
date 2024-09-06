package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Setting struct {
	User     string
	Host     string
	Password string
	Port     string
	DBName   string
}

func ReadEnv() *Setting {
	var app Setting

	err := godotenv.Load(".env");

	if err != nil {
		return &Setting{}
	};

	app.User = os.Getenv("DB_USER");
	app.Host = os.Getenv("DB_HOST");
	app.Port = os.Getenv("DB_PORT");
	app.DBName = os.Getenv("DB_NAME");
	app.Password = os.Getenv("DB_PASSWORD")
	
	return &app
};


func InitConfig() *Setting {
	return ReadEnv();
}