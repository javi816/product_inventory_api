package dto

type CreateProductDTO struct {
	Name        string
	Description string
	Price       float64
	Stock       int
}

type UpdateProductDTO struct {
	Name        *string
	Description *string
	Price       *float64
	Stock       *int
}

type ProductResponseDTO struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Stock       int
}
