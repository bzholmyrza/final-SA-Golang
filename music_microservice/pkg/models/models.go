package models

import (
	"errors"
)

var ErrNoRecord = errors.New("models: no matching record found")

type Song struct {
	ID          int
	Title       string
	Author      string
	ReleaseDate string
}
