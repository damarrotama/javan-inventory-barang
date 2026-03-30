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
	ProductController      controller.ProductController
	StockController        controller.StockController
	StockHistoryController controller.StockHistoryController
}

type Domain struct {
	ProductDomain      domain.ProductDomain
	StockDomain        domain.StockDomain
	StockHistoryDomain domain.StockHistoryDomain
}

type Repository struct {
	ProductRepository      repository.ProductRepository
	StockRepository        repository.StockRepository
	StockHistoryRepository repository.StockHistoryRepository
}

func NewService() *Service {
	return &Service{}
}
