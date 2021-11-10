package repository

import "go-learn-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
