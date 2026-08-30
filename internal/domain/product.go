package domain

import "errors"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
}

func NewProduct(name, description string, price float64, stock int) (*Product, error) {

	if name == "" {
		return nil, errors.New("Product name is empty")
	}

	if price <= 0 {
		return nil, errors.New("Wrong Product price")
	}

	if stock < 0 {
		return nil, errors.New("Wrong Product stock")
	}

	product := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	return product, nil
}
