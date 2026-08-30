package ports

import "github.com/javi816/product_inventory_api/internal/domain"

type ProductRepository interface {
	Create(product *domain.Product) (*domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	GetAll() ([]*domain.Product, error)
	Update(product *domain.Product) (*domain.Product, error)
	Delete(id int) error
}
