package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"musinfo/internal/handlers"
	"musinfo/internal/repository"
)

func Song(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	repo := repository.NewSongRepository()
	handler := handlers.NewSongHandler(*repo)

	e.POST("/song", handler.CreateSong)
	e.PUT("/song/:id", handler.UpdateSong)
	e.GET("/song/:id", handler.GetSong)
	e.GET("/song", handler.ListSongs)
	e.DELETE("/song/:id", handler.DeleteSong)
}
