package main

import (
	"twintask/apps/config"
	"twintask/apps/database"
	"twintask/apps/migrations"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New();
    
	cfg := config.InitConfig();
	DBPosgresql := database.InitDB(cfg);
	migrations.InitMigration(DBPosgresql);

	e.Pre(middleware.RemoveTrailingSlash());
	e.Use(middleware.Logger());

	e.Use(middleware.CORS());

	e.Logger.Fatal(e.Start(":9999"));
}