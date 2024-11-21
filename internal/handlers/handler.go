package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"musinfo/internal/models"
	"musinfo/internal/repository"
	"musinfo/internal/responses"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type SongHandler struct {
	Repo repository.SongRepository
}

func NewSongHandler(repo repository.SongRepository) *SongHandler {
	return &SongHandler{Repo: repo}
}

func (h *SongHandler) CreateSong(c echo.Context) error {
	var input models.Song
	if err := c.Bind(&input); err != nil {
		return responses.BadRequestResponse(c, err)
	}
	url := fmt.Sprintf(
		"%s/info?group=%s&song=%s",
		os.Getenv("EXTERNAL_API_URL"),
		url.QueryEscape(input.Group),
		url.QueryEscape(input.Song),
	)

	externalApiResponse, err := http.Get(url)
	if err != nil {
		return responses.InternalServerErrorResponse(c, fmt.Errorf("Failed to connect to %s: %w", os.Getenv("EXTERNAL_API_URL"), err))
	}
	var closeErr error
	defer func() {
		if err := externalApiResponse.Body.Close(); err != nil {
			closeErr = fmt.Errorf("Failed to close response body: %w", err)
		}
	}()

	if closeErr != nil {
		return responses.InternalServerErrorResponse(c, closeErr)
	}

	if externalApiResponse.StatusCode != http.StatusOK {
		return responses.InternalServerErrorResponse(c, fmt.Errorf("External API Error: %w", err))
	}

	var additionalInfo models.Song
	if err := json.NewDecoder(externalApiResponse.Body).Decode(&additionalInfo); err != nil {
		return responses.InternalServerErrorResponse(c, fmt.Errorf("JSON Decode Error: %w", err))
	}

	validate := validator.New()

	if err := validate.Struct(additionalInfo); err != nil {
		return responses.InternalServerErrorResponse(c, fmt.Errorf("Validate data error: %w", err))
	}

	input.ReleaseDate = additionalInfo.ReleaseDate
	input.Text = additionalInfo.Text
	input.Link = additionalInfo.Link

	result, err := h.Repo.Create(input)
	if err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}
	return responses.SuccessResponse(c, result)
}

func (h *SongHandler) GetSong(c echo.Context) error {
	parsedUUID, err := uuid.Parse(c.Param("id"))
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	if err != nil {
		return responses.BadRequestResponse(c, err)
	}
	pagination := models.Pagination{}
	page, err := strconv.Atoi(pageParam)
	if err == nil {
		pagination.Page = page
	}

	limit, err := strconv.Atoi(limitParam)
	if err == nil {
		pagination.Limit = limit
	}
	result, err := h.Repo.GetByIdWithPagination(parsedUUID, pagination)
	if err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}
	return responses.SuccessResponse(c, result)
}

func (h *SongHandler) ListSongs(c echo.Context) error {
	var input models.Song

	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")

	input.Group = c.QueryParam("group")
	input.Song = c.QueryParam("song")
	input.ReleaseDate = c.QueryParam("releaseDate")
	input.Link = c.QueryParam("link")
	filterText := c.QueryParam("text")
	pagination := models.Pagination{}
	page, err := strconv.Atoi(pageParam)
	if err == nil {
		pagination.Page = page
	}

	limit, err := strconv.Atoi(limitParam)
	if err == nil {
		pagination.Limit = limit
	}

	res, err := h.Repo.List(input, pagination, filterText)
	if err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}
	return responses.SuccessResponse(c, res)
}

func (h *SongHandler) UpdateSong(c echo.Context) error {
	parsedUUID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return responses.BadRequestResponse(c, err)
	}

	var input models.Song
	if err := c.Bind(&input); err != nil {
		return responses.BadRequestResponse(c, err)
	}

	if err := h.Repo.Update(parsedUUID, input); err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}
	return responses.NoContentResponse(c)
}

func (h *SongHandler) DeleteSong(c echo.Context) error {
	parsedUUID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return responses.BadRequestResponse(c, err)
	}
	if err := h.Repo.Delete(parsedUUID); err != nil {
		return responses.InternalServerErrorResponse(c, err)
	}
	return responses.NoContentResponse(c)
}
