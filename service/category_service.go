package service

import (
	"errors"

	"github.com/fatoni-ach/app-sayhello/entities"
	"github.com/fatoni-ach/app-sayhello/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entities.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("Category Not Found")
	}

	return category, nil
}
