package postgresql

import (
	"context"
	"errors"
	"final-SA-Golang/favorites_microservice/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type FavoritesModel struct {
	DB *pgxpool.Pool
}

func (m *FavoritesModel) CreateFavorites(UserID, SongID int) (int, error) {
	stmt := `INSERT INTO favorites (user_id, song_id)
			VALUES($1, $2) RETURNING id`

	var id int
	err := m.DB.QueryRow(context.Background(), stmt, UserID, SongID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *FavoritesModel) DeleteFavorites(UserID, SongID int) (*models.Favorites, error) {
	stmt := `DELETE FROM favorites
			WHERE user_id = $1 AND song_id = $2`

	row := m.DB.QueryRow(context.Background(), stmt, UserID, SongID)
	s := &models.Favorites{}

	err := row.Scan(&s.ID, &s.UserID, &s.SongID)
	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *FavoritesModel) GetFavoritesByUserID(UserID int) ([]*models.Favorites, error) {
	stmt := `SELECT song_id FROM favorites
			 WHERE user_id = $1`

	rows, err := m.DB.Query(context.Background(), stmt, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	favorites := []*models.Favorites{}

	for rows.Next() {
		s := &models.Favorites{}
		err = rows.Scan(&s.ID, &s.UserID, &s.SongID)
		if err != nil {
			return nil, err
		}

		favorites = append(favorites, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return favorites, nil

}
