package postgresql

import (
	"context"
	"errors"
	"final-SA-Golang/music_microservice/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SongModel struct {
	DB *pgxpool.Pool
}

func (m *SongModel) CreateSong(title, author, releaseDate string) (int, error) {
	stmt := `INSERT INTO songs (title, author, release_date)
			VALUES($1, $2, $3) RETURNING id`

	var id int
	err := m.DB.QueryRow(context.Background(), stmt, title, author, releaseDate).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *SongModel) UpdateSong(title, author, releaseDate string, id int) (int, error) {
	stmt := `UPDATE songs SET title = $1, author = $2, 
        release_date = $3
          WHERE id = $4 RETURNING id`

	err := m.DB.QueryRow(context.Background(), stmt, title, author, releaseDate, id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *SongModel) DeleteSong(id int) (*models.Song, error) {
	stmt := `DELETE FROM songs
			WHERE id = $1`

	row := m.DB.QueryRow(context.Background(), stmt, id)
	s := &models.Song{}

	err := row.Scan(&s.ID, &s.Title, &s.Author, &s.ReleaseDate)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SongModel) GetSongByID(id int) (*models.Song, error) {
	stmt := `SELECT id, title, author, release_date FROM songs
			WHERE id = $1`

	row := m.DB.QueryRow(context.Background(), stmt, id)
	s := &models.Song{}

	err := row.Scan(&s.ID, &s.Title, &s.Author, &s.ReleaseDate)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}
func (m *SongModel) GetAllSongs() ([]*models.Song, error) {
	stmt := `SELECT id, title, author, release_date FROM songs
			 ORDER BY title `

	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := []*models.Song{}

	for rows.Next() {
		s := &models.Song{}
		err = rows.Scan(&s.ID, &s.Title, &s.Author, &s.ReleaseDate)
		if err != nil {
			return nil, err
		}

		songs = append(songs, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil

}
