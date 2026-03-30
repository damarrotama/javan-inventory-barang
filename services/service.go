package services

import (
	"javan-inventory-barang/controller"
	"javan-inventory-barang/domain"
	"javan-inventory-barang/repository"
)

type Service struct {
	Controller Controller
	Domain     Domain
	Repository Repository
}

type Controller struct {
	ProductController controller.ProductController
}

type Domain struct {
	ProductDomain domain.ProductDomain
}

type Repository struct {
	ProductRepository repository.ProductRepository
}

func NewService() *Service {
	return &Service{}
}
