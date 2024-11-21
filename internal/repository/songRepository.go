package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"musinfo/internal/db"
	"musinfo/internal/models"
	"strconv"
	"strings"
)

type SongRepository struct{}

func NewSongRepository() *SongRepository {
	return &SongRepository{}
}

func (s *SongRepository) Create(songInfo models.Song) (uuid.UUID, error) {

	query := `INSERT INTO songs (group_name, song, release_date, song_text, link) 
              VALUES ($1, $2, $3, $4, $5)
              RETURNING id`
	var id string
	err := db.DB.QueryRow(
		query,
		songInfo.Group,
		songInfo.Song,
		songInfo.ReleaseDate,
		songInfo.Text,
		songInfo.Link,
	).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, err
	}
	return parsedUUID, nil
}

func (s *SongRepository) List(filter models.Song, pagination models.Pagination, text string) ([]models.Song, error) {
	var songs []models.Song
	query := `SELECT group_name, song, release_date, song_text, link FROM songs `
	fieldNames, values := generateQueryFields(filter)
	if len(fieldNames) > 0 {
		query += " WHERE " + strings.Join(fieldNames, " AND ")
	}
	if len(text) > 0 {
		if len(values) > 0 {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " song_text ILIKE $" + strconv.Itoa(len(values)+1) + ""
		values = append(values, "%"+text+"%")
	}
	if pagination.Limit > 0 {
		query += " LIMIT $" + strconv.Itoa(len(values)+1)
		values = append(values, pagination.Limit)
	}
	if pagination.Page > 0 {
		query += " OFFSET $" + strconv.Itoa(len(values)+1)
		values = append(values, pagination.Page)
	}
	rows, err := db.DB.Query(query, values...)
	if err != nil {
		return []models.Song{}, fmt.Errorf("Error in query: %s . Detail: %w", query, err)
	}
	defer rows.Close()
	for rows.Next() {
		var song models.Song
		err := rows.Scan(
			&song.Group,
			&song.Song,
			&song.ReleaseDate,
			&song.Text,
			&song.Link,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		songs = append(songs, song)
	}
	fmt.Printf("Query: %v\n", query)
	return songs, nil
}

func (s *SongRepository) GetByIdWithPagination(ID uuid.UUID, pagination models.Pagination) (models.Song, error) {
	query := `SELECT group_name, song, song_text FROM songs where id = $1`

	// Выполнение запроса
	var song models.Song
	err := db.DB.QueryRow(query, ID).Scan(
		&song.Group,
		&song.Song,
		&song.Text,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Song{}, fmt.Errorf("song with ID %s not found", ID)
		}
		return models.Song{}, fmt.Errorf("song %s with ID %w", query, err)
	}
	song.Text = parseText(song.Text, pagination)

	return song, nil
}

func (s *SongRepository) Update(ID uuid.UUID, song models.Song) error {
	query := `UPDATE songs SET `
	fieldNames, values := generateQueryFields(song)

	if len(fieldNames) == 0 {
		return fmt.Errorf("no valid fields to update")
	}

	query += strings.Join(fieldNames, ", ") + " WHERE id = $" + strconv.Itoa(len(values)+1)

	values = append(values, ID)

	if _, err := db.DB.Exec(query, values...); err != nil {
		return err
	}

	return nil
}

func (s *SongRepository) Delete(ID uuid.UUID) error {
	query := `DELETE FROM songs WHERE id = $1`
	_, err := db.DB.Exec(query, ID.String())
	if err != nil {
		return err
	}
	return nil
}

func parseText(text string, pagination models.Pagination) string {
	if pagination.Page <= 0 || pagination.Limit <= 0 {
		return text
	}
	songParts := strings.Split(text, "\n")
	var trimParts []string
	for _, str := range songParts {
		if str != "" {
			trimParts = append(trimParts, str)
		}
	}

	startIndex := (pagination.Page - 1) * pagination.Limit
	endIndex := pagination.Page * pagination.Limit

	if endIndex > len(trimParts) {
		endIndex = len(trimParts)
	}
	if startIndex > len(trimParts) {
		return ""
	}

	return strings.Join(trimParts[startIndex:endIndex], "\n")
}

func generateQueryFields(song models.Song) ([]string, []interface{}) {
	var values []interface{}
	var fieldNames []string

	// Для обновления данных проверяем каждое поле и добавляем его в запрос, если оно передано
	if song.Group != "" {
		fieldNames = append(fieldNames, "group_name = $"+strconv.Itoa(len(values)+1))
		values = append(values, song.Group)
	}
	if song.Song != "" {
		fieldNames = append(fieldNames, "song = $"+strconv.Itoa(len(values)+1))
		values = append(values, song.Song)
	}
	if song.ReleaseDate != "" {
		fieldNames = append(fieldNames, "release_date = $"+strconv.Itoa(len(values)+1))
		values = append(values, song.ReleaseDate)
	}
	if song.Text != "" {
		fieldNames = append(fieldNames, "song_text = $"+strconv.Itoa(len(values)+1))
		values = append(values, song.Text)
	}
	if song.Link != "" {
		fieldNames = append(fieldNames, "link = $"+strconv.Itoa(len(values)+1))
		values = append(values, song.Link)
	}
	return fieldNames, values
}
