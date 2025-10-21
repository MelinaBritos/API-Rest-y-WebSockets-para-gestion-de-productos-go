package repository

import "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"

type ProductCategoryRepository interface {
	CreateProductCategory(productID int, categoryID int) error
	GetProductCategories(productID int) ([]*Model.Category, error)
	DeleteProductCategory(productID int, categoryID int) error
}
