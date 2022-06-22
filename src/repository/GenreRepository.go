package repository

import (
	"gorm.io/gorm"
	"sakaba.link/api/src/model"
)

// GenreRepository is responsible for reading from and writing to DB Table `genres`.
type GenreRepository struct {
	DB *gorm.DB
}

// GetAllGenres returns all the genres.
func (c GenreRepository) GetAllGenres() []model.Genre {
	allGenres := []model.Genre{}
	c.DB.Table("genres").
		Select("id", "name").
		Order("name ASC").
		Scan(&allGenres)

	return allGenres
}

// GetGenreByID returns the specified genre.
func (c GenreRepository) GetGenreByID(id string) model.Genre {
	genre := model.Genre{}
	c.DB.Table("genres").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&genre)

	return genre
}
