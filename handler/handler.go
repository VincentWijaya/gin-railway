package handler

import "gin-railway/service"

type Handler struct {
	productService *service.ProductService
}

func NewHandler(productService *service.ProductService) *Handler {
	return &Handler{
		productService: productService,
	}
}
