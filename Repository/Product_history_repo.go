package repository

import "github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"

type ProductHistoryRepository interface {
	CreateProductHistory(history Model.ProductHistory) error
	GetProductHistoryByID(id int) (*Model.ProductHistory, error)
	UpdateProductHistory(history Model.ProductHistory) error
	DeleteProductHistory(id int) error
	GetAllProductHistories() ([]*Model.ProductHistory, error)
}
