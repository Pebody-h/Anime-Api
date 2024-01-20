package main

import (
	"github.com/Pebody-h/Anime-Api/internal/database"
	"github.com/Pebody-h/Anime-Api/routes"
	"github.com/Pebody-h/Anime-Api/settings"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := settings.Settings("PORT")
	if port == "" {
		port = "8080"
	}

	database.ConnectDB()

	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.SetupRoutes(app)

	app.Logger.Fatal(app.Start(":" + port))
}
