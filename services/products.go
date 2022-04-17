package services

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Skiu        string
	Price       float64
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

var ProductList = []*Product{
	&Product{
		ID:          1,
		Name:        "Coffee",
		Description: "Chocolate and creamy  beverage",
		Skiu:        "adc123",
		Price:       35.00,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Latte",
		Description: "Milky, very creamy beverage",
		Skiu:        "adc234",
		Price:       50.00,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
