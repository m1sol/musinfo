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

// Documentation comments for Swagger

// @Router /song [post]
// @Summary Добавление новой песни
// @Description Добавление новой песни (название, группа, )
// @Tags songs
// @Accept json
// @Produce json
// @Param song body models.CreateSong true "Song details"
// @Success 200 {object} responses.Response{data=models.Song}
// @Failure 500 {object} responses.Response
func createSongDocs() {}

// @Router /song/{id} [put]
// @Summary Update an existing song
// @Description Update song details by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param song body models.Song true "Updated song details"
// @Success 200 {object} responses.Response{data=models.Song} "Song updated successfully"
// @Failure 404 {object} responses.Response "Song not found"
func updateSongDocs() {}

// @Router /song/{id} [get]
// @Summary Get a song by ID
// @Description Retrieve song details by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} responses.Response{data=models.OutputSong} "Song retrieved successfully"
// @Failure 404 {object} responses.Response "Song not found"
func getSongDocs() {}

// @Router /song [get]
// @Summary List all songs
// @Description Retrieve a list of all songs
// @Tags songs
// @Accept json
// @Produce json
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Param song query string false "Song Name"
// @Param group query string false "Group Name"
// @Param releaseDate query string false "Release Date (DD.MM.YYYY)"
// @Param link query string false "Link"
// @Param text query string false "Text (part of song)"
// @Success 200 {object} responses.Response{data=[]models.OutputSong} "List of songs"
// @Failure 500 {object} responses.Response
func listSongDocs() {}

// @Router /song/{id} [delete]
// @Summary Delete a song by ID
// @Description Remove a song from the system by ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path string true "Song ID"
// @Success 204 {object} responses.Response "Song deleted successfully"
// @Failure 404 {object} responses.Response "Song not found"
func deleteSongDocs() {}
