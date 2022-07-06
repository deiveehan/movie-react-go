package models

import (
	"database/sql"
	"time"
)

// Models is wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns model with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Movie type for movies
type Movie struct {
	ID          int          `json: "id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"release_date"`
	Runtime     int          `json:"runtime"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	MovieGenre  []MovieGenre `json:"-"`
}

// Genre type for genre
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// MovieGenre type for Moviegenre

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movie_id"`
	GenreID   int       `json:"genre_id"`
	Genre     int       `json:"genre"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
