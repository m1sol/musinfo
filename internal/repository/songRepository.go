package repository

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"musinfo/internal/db"
	"musinfo/internal/entities"
	"musinfo/internal/structs"
	"reflect"
)

type SongRepository struct{}

func NewSongRepository() *SongRepository {
	return &SongRepository{}
}

func (s *SongRepository) Create(songInfo entities.Song) (uuid.UUID, error) {

	query := `INSERT INTO songs ("group", song, release_date, text, link) 
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

func (s *SongRepository) List(filter entities.Song, pagination structs.Pagination) (any /*[]entities.Song*/, error) {
	v := reflect.ValueOf(filter)
	typeOfFilter := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := typeOfFilter.Field(i)
		fmt.Printf("filter: %+v\n", fieldType.Tag.Get("json"))
		fmt.Printf("val: %+v\n", field)
	}
	fmt.Printf("filter: %+v\n", v)
	return v /*[]entities.Song{}*/, nil
}

func (s *SongRepository) GetByIdWithPagination(ID uuid.UUID, pagination structs.Pagination) (entities.Song, error) {
	query := `SELECT "group", song, text FROM songs where id = $1`

	// Выполнение запроса
	var song entities.Song
	err := db.DB.QueryRow(query, ID).Scan(
		ID.String(),
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return entities.Song{}, fmt.Errorf("song with ID %s not found", ID)
		}
		return entities.Song{}, fmt.Errorf("song %s with ID %w", query, err)
	}

	return song, nil

}

func (s *SongRepository) Update(ID uuid.UUID, entity entities.Song) error {
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
