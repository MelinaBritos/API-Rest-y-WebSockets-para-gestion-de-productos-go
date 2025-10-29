package Repository

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
)

type ValidFilters struct {
	Search   string
	Page     int
	Limit    int
	SortBy   string
	Order    string
	MinPrice float64
	MaxPrice float64
}

type SearchRepository interface {
	SearchProducts(filters ValidFilters) ([]*Model.Product, error)
	SearchCategories(filters ValidFilters) ([]*Model.Category, error)
}

type searchRepository struct {
	db *sql.DB
}

func NewSearchRepository(db *sql.DB) SearchRepository {
	return &searchRepository{db: db}
}

func (r *searchRepository) SearchProducts(filters ValidFilters) ([]*Model.Product, error) {
	baseQuery := "SELECT id, name, description, price, stock FROM products"

	args := []interface{}{}
	paramCount := 1
	whereClauses := []string{}

	if filters.Search != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("(name ILIKE $%d)", paramCount))
		args = append(args, "%"+filters.Search+"%")
		paramCount++
	}

	if filters.MinPrice > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("price >= $%d", paramCount))
		args = append(args, filters.MinPrice)
		paramCount++
	}

	if filters.MaxPrice > 0 {
		whereClauses = append(whereClauses, fmt.Sprintf("price <= $%d", paramCount))
		args = append(args, filters.MaxPrice)
		paramCount++
	}

	fullQuery := baseQuery
	if len(whereClauses) > 0 {
		fullQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	fullQuery += fmt.Sprintf(" ORDER BY %s %s", filters.SortBy, filters.Order)

	// Numero de paginas a saltar usando el limite de productos ingresado por pagina
	offset := (filters.Page - 1) * filters.Limit
	fullQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramCount, paramCount+1)
	args = append(args, filters.Limit, offset)

	rows, err := r.db.Query(fullQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*Model.Product{}
	for rows.Next() {
		product := Model.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (r *searchRepository) SearchCategories(filters ValidFilters) ([]*Model.Category, error) {
	baseQuery := "SELECT id, name, description FROM categories"

	args := []interface{}{}
	paramCount := 1
	whereClauses := []string{}

	if filters.Search != "" {
		whereClauses = append(whereClauses, fmt.Sprintf("(name ILIKE $%d)", paramCount))
		args = append(args, "%"+filters.Search+"%")
		paramCount++
	}

	fullQuery := baseQuery
	if len(whereClauses) > 0 {
		fullQuery += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	fullQuery += fmt.Sprintf(" ORDER BY %s %s", filters.SortBy, filters.Order)

	// Numero de paginas a saltar usando el limite de categorias ingresado
	offset := (filters.Page - 1) * filters.Limit
	fullQuery += fmt.Sprintf(" LIMIT $%d OFFSET $%d", paramCount, paramCount+1)
	args = append(args, filters.Limit, offset)

	rows, err := r.db.Query(fullQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []*Model.Category{}
	for rows.Next() {
		category := Model.Category{}
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
