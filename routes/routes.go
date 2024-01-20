package routes

import (
	"github.com/Pebody-h/Anime-Api/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	app.GET("/animes", handlers.GetAnimes)
	app.GET("/anime/:id", handlers.GetAnimeById)
	app.POST("/animes", handlers.CreateAnimes)
	app.POST("/anime", handlers.CreateAnime)
	app.PUT("/anime/:id", handlers.UpdateAnime)
	app.DELETE("/anime/:id", handlers.DeleteAnime)
}
