package handlers

import (
	"net/http"

	"github.com/Pebody-h/Anime-Api/internal/database"
	"github.com/Pebody-h/Anime-Api/models"
	"github.com/labstack/echo/v4"
)

func GetAnimes(c echo.Context) error {
	var animes []models.Animes

	if err := database.DB.Find(&animes).Error; err != nil {
		// Manejar el error (log, respuesta HTTP, etc.)
		return c.JSON(http.StatusInternalServerError, "Error obteniendo animes")
	}

	return c.JSON(http.StatusOK, animes)
}

func GetAnimeById(c echo.Context) error {
	id := c.Param("id")
	var anime models.Animes

	if err := database.DB.First(&anime, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Anime not found")
	}

	return c.JSON(http.StatusOK, anime)
}

func CreateAnimes(c echo.Context) error {
	var animes []models.Animes

	if err := c.Bind(&animes); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	for _, anime := range animes {
		database.DB.Create(&anime)
	}

	return c.JSON(http.StatusCreated, animes)
}

func CreateAnime(c echo.Context) error {
	anime := new(models.Animes)

	if err := c.Bind(anime); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	database.DB.Create(&anime)

	return c.JSON(http.StatusCreated, anime)
}

func UpdateAnime(c echo.Context) error {
	id := c.Param("id")
	var anime models.Animes

	if err := database.DB.First(&anime, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Anime not found")
	}

	if err := c.Bind(&anime); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	database.DB.Save(&anime)

	return c.JSON(http.StatusOK, anime)
}

func DeleteAnime(c echo.Context) error {
	id := c.Param("id")
	var anime models.Animes

	if err := database.DB.First(&anime, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Todo not found")
	}

	database.DB.Delete(&anime)

	return c.NoContent(http.StatusNoContent)
}
