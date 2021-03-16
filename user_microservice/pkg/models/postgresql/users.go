package postgresql

import (
	"context"
	"errors"
	"final-SA-Golang/user_microservice/pkg/models"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserModel struct {
	DB *pgxpool.Pool
}

func (m *UserModel) CreateUser(username, email, password string, role int) (int, error) {
	stmt := `INSERT INTO users (username, email, password, role)
      VALUES($1, $2, $3, $4) RETURNING id`
	var id int
	err := m.DB.QueryRow(context.Background(), stmt, username, email, password, role).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *UserModel) UpdateUser(username, email, password string, id, role int) (int, error) {
	stmt := `UPDATE users SET username = $1, email = $2, 
				password = $3, role = $4 WHERE id = $5 RETURNING id`
	var id2 int
	err := m.DB.QueryRow(context.Background(), stmt, username, email, password, role, id).Scan(&id2)
	if err != nil {
		return 0, err
	}
	return id2, nil
}

func (m *UserModel) DeleteUser(id int) pgx.Row {
	stmt := `DELETE FROM users WHERE id = $1`
	err := m.DB.QueryRow(context.Background(), stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) GetUser(id int) (*models.User, error) {
	stmt := `SELECT id, username, email, role, password FROM users
      		WHERE id = $1`

	row := m.DB.QueryRow(context.Background(), stmt, id)
	s := &models.User{}

	err := row.Scan(&s.ID, &s.Username, &s.Email, &s.Role, &s.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *UserModel) GetUserByEmailAndPassword(email, password string) (*models.User, error) {
	stmt := `SELECT id, username, email, role, password FROM users
      		WHERE email = $1 AND password = $2`

	row := m.DB.QueryRow(context.Background(), stmt, email, password)
	s := &models.User{}

	err := row.Scan(&s.ID, &s.Username, &s.Email, &s.Role, &s.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

func (m *UserModel) GetAllUsers() ([]*models.User, error) {
	stmt := `SELECT id, username, email, password, role FROM users
			ORDER BY id`

	rows, err := m.DB.Query(context.Background(), stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []*models.User{}

	for rows.Next() {
		s := &models.User{}
		err = rows.Scan(&s.ID, &s.Username, &s.Email, &s.Password, &s.Role)
		if err != nil {
			return nil, err
		}

		users = append(users, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
