package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// GenreRepository is responsible for reading from and writing to DB Table `genres`.
type GenreRepository struct{}

// GetAllGenres returns all the genres.
func (c *GenreRepository) GetAllGenres() []model.Genre {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	allGenres := []model.Genre{}
	db.Table("genres").
		Select("id", "name").
		Order("name ASC").
		Scan(&allGenres)

	return allGenres
}

// GetGenreByID returns the specified genre.
func (c *GenreRepository) GetGenreByID(id string) model.Genre {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	genre := model.Genre{}
	db.Table("genres").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&genre)

	return genre
}
