package models

import (
	"github.com/google/uuid"
	"time"
)

type Song struct {
	Group       string `json:"group"`
	Song        string `json:"song"`
	ReleaseDate string `json:"releaseDate" validate:"required"`
	Text        string `json:"text" validate:"required"`
	Link        string `json:"link" validate:"required"`
}

type CreateSong struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type OutputSong struct {
	ID          uuid.UUID `json:"id"`
	Group       string    `json:"group"`
	Song        string    `json:"song"`
	ReleaseDate string    `json:"releaseDate" validate:"required"`
	Text        string    `json:"text" validate:"required"`
	Link        string    `json:"link" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
