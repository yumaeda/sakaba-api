package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// GenreRepository is responsible for reading from and writing to DB Table `genres`.
type GenreRepository struct{}

// GetAllGenres returns all the genres.
func (c *GenreRepository) GetAllGenres() []model.Genre {
	allGenres := []model.Genre{}
	db := infrastructure.ConnectToDB()
	db.Table("genres").
		Select("id", "name").
		Order("name ASC").
		Scan(&allGenres)
	infrastructure.CloseDB(db)

	return allGenres
}

// GetGenreByID returns the specified genre.
func (c *GenreRepository) GetGenreByID(id string) model.Genre {
	genre := model.Genre{}
	db := infrastructure.ConnectToDB()
	db.Table("genres").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&genre)
	infrastructure.CloseDB(db)

	return genre
}
