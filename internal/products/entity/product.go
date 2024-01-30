package entity

import "github.com/google/uuid"

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	Category    *Category
	ImageURL    string
}

func NewProduct(name string, description string, price float64, category *Category, imageURL string) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
	}
}
