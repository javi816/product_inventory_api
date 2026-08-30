package usecases

import (
	"errors"

	"github.com/javi816/product_inventory_api/internal/application/dto"
	"github.com/javi816/product_inventory_api/internal/application/ports"
	"github.com/javi816/product_inventory_api/internal/domain"
)

type ProductUseCase struct {
	repository ports.ProductRepository
}

func (puc ProductUseCase) CreateProduct(product dto.CreateProductDTO) (*dto.ProductResponseDTO, error) {

	newProduct, err := domain.NewProduct(
		product.Name,
		product.Description,
		product.Price,
		product.Stock,
	)

	if err != nil {
		return nil, err
	}

	savedProduct, err := puc.repository.Create(newProduct)

	if err != nil {
		return nil, err
	}

	mappedSavedProduct := &dto.ProductResponseDTO{
		ID:          savedProduct.ID,
		Name:        savedProduct.Name,
		Description: savedProduct.Description,
		Price:       savedProduct.Price,
		Stock:       savedProduct.Stock,
	}

	return mappedSavedProduct, nil
}

func (puc ProductUseCase) GetProduct(id int) (*dto.ProductResponseDTO, error) {
	product, err := puc.getByID(id)

	if err != nil {
		return nil, err
	}

	productResponse := &dto.ProductResponseDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	return productResponse, nil
}

func (puc ProductUseCase) ListProducts() ([]dto.ProductResponseDTO, error) {

	products, err := puc.repository.GetAll()

	if err != nil {
		return nil, err
	}

	productsSlice := make([]dto.ProductResponseDTO, 0, len(products))

	for _, product := range products {
		mappedProduct := dto.ProductResponseDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		}

		productsSlice = append(productsSlice, mappedProduct)
	}

	return productsSlice, nil

}

func (puc ProductUseCase) UpdateProduct(id int, productData *dto.UpdateProductDTO) (*dto.ProductResponseDTO, error) {
	product, err := puc.getByID(id)

	if err != nil {
		return nil, err
	}

	if productData.Name != nil {
		product.Name = *productData.Name
	}

	if productData.Description != nil {
		product.Description = *productData.Description
	}

	if productData.Price != nil {
		product.Price = *productData.Price
	}

	if productData.Stock != nil {
		product.Stock = *productData.Stock
	}

	updatedProduct, err := puc.repository.Update(product)

	if err != nil {
		return nil, err
	}

	responseProduct := &dto.ProductResponseDTO{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		Stock:       updatedProduct.Stock,
	}

	return responseProduct, nil
}

func (puc ProductUseCase) DeleteProduct(id int) error {
	_, err := puc.getByID(id)

	if err != nil {
		return err
	}

	if err = puc.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (puc ProductUseCase) getByID(id int) (*domain.Product, error) {
	product, err := puc.repository.GetByID(id)

	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}
