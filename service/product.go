package service

import "gin-railway/repository"

type ProductService struct {
	productRepo repository.ProductRepo
}

func NewService(productRepo repository.ProductRepo) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}
