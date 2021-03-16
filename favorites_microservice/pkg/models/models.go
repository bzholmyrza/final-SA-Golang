package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Favorites struct {
	ID     int
	UserID int
	SongID int
}
