package repository

import "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"

type CategoryRepository interface {
	CreateCategory(category Model.Category) error
	GetCategoryByID(id int) (*Model.Category, error)
	UpdateCategory(category Model.Category) error
	DeleteCategory(id int) error
	GetAllCategories() ([]*Model.Category, error)
}
