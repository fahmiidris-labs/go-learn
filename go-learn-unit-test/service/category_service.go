package service

import (
	"errors"
	"go-learn-unit-test/entity"
	"go-learn-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Err")
	} else {
		return category, nil
	}
}
