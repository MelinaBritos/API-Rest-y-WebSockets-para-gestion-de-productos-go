package Service

import (
	"errors"
	"strconv"

	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Model"
	"github.com/MelinaBritos/API-REST-y-WebSockets-para-gestion-de-productos/Repository"
)

type SearchService struct {
	repo Repository.SearchRepository
}

func NewSearchService(repo Repository.SearchRepository) *SearchService {
	return &SearchService{repo: repo}
}

func (s *SearchService) SearchProducts(filters map[string]string) ([]*Model.Product, error) {

	page, _ := strconv.Atoi(filters["page"])
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(filters["limit"])
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	order := filters["order"]
	if order == "" {
		order = "ASC"
	}
	if order != "ASC" && order != "DESC" {
		return nil, errors.New("el orden debe ser ASC o DESC")
	}

	sortBy := filters["sort_by"]
	if sortBy == "" {
		sortBy = "id"
	}
	if sortBy != "name" && sortBy != "price" && sortBy != "stock" && sortBy != "id" {
		return nil, errors.New("el campo de ordenación no es válido")
	}

	minPrice, _ := strconv.ParseFloat(filters["min_price"], 64)
	maxPrice, _ := strconv.ParseFloat(filters["max_price"], 64)

	validFilters := Repository.ValidFilters{
		Search:   filters["search"],
		MinPrice: minPrice,
		MaxPrice: maxPrice,
		SortBy:   sortBy,
		Order:    order,
		Page:     page,
		Limit:    limit,
	}

	return s.repo.SearchProducts(validFilters)
}

func (s *SearchService) SearchCategories(filters map[string]string) ([]*Model.Category, error) {

	page, _ := strconv.Atoi(filters["page"])
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(filters["limit"])
	if limit <= 0 || limit > 50 {
		limit = 10
	}

	order := filters["order"]
	if order == "" {
		order = "ASC"
	}
	if order != "ASC" && order != "DESC" {
		return nil, errors.New("el orden debe ser ASC o DESC")
	}

	sortBy := filters["sort_by"]
	if sortBy == "" {
		sortBy = "id"
	}
	if sortBy != "id" {
		return nil, errors.New("el campo de ordenación no es válido")
	}

	validFilters := Repository.ValidFilters{
		Search: filters["search"],
		Page:   page,
		Limit:  limit,
		SortBy: sortBy,
		Order:  order,
	}

	return s.repo.SearchCategories(validFilters)
}
