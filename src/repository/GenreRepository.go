package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

type GenreRepository struct{}

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

func (c *GenreRepository) GetGenreById(id string) model.Genre {
	genre := model.Genre{}
	db := infrastructure.ConnectToDB()
	db.Table("genres").
		Select("id", "name").
		Where("id = ?", id).
		Scan(&genre)
	infrastructure.CloseDB(db)

	return genre
}
