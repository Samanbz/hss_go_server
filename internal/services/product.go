package services

import (
	"context"
	"hss/internal/models"
	"hss/internal/repositories"
)

type ProductService struct {
	productRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) InsertProduct(ctx context.Context, product *models.Product) error {
	err := s.productRepository.Create(context.Background(), product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	product, err := s.productRepository.GetByID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
