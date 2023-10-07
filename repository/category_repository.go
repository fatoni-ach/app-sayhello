package repository

import "github.com/fatoni-ach/app-sayhello/entities"

type CategoryRepository interface {
	FindById(id string) *entities.Category
}
