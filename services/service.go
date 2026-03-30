package services

import (
	"javan-inventory-barang/controller"
	"javan-inventory-barang/domain"
	"javan-inventory-barang/repository"
	"javan-inventory-barang/transaction"
	"javan-inventory-barang/utils"
	"log"
)

type Service struct {
	Controller Controller
	Domain     Domain
	Repository Repository
}

type Controller struct {
	ProductController controller.ProductController
	StockController   controller.StockController
}

type Domain struct {
	ProductDomain domain.ProductDomain
	StockDomain   domain.StockDomain
}

type Repository struct {
	ProductRepository      repository.ProductRepository
	StockRepository        repository.StockRepository
	StockHistoryRepository repository.StockHistoryRepository
}

// NewService wires repositories, domain, and controllers to a shared DB pool.
func NewService() *Service {
	db, err := utils.OpenPostgres()
	if err != nil {
		log.Fatal(err)
	}

	txManager := transaction.NewManager(db)

	productRepo := repository.NewProductRepository(db)
	stockRepo := repository.NewStockRepository(db)
	stockHistoryRepo := repository.NewStockHistoryRepository(db)

	productDomain := domain.NewProductDomain(productRepo)
	stockDomain := domain.NewStockDomain(stockRepo, stockHistoryRepo, productRepo, txManager)

	productController := controller.NewProductController(productDomain)
	stockController := controller.NewStockController(stockDomain)

	return &Service{
		Controller: Controller{
			ProductController: productController,
			StockController:   stockController,
		},
		Domain: Domain{
			ProductDomain: productDomain,
			StockDomain:   stockDomain,
		},
		Repository: Repository{
			ProductRepository:      productRepo,
			StockRepository:        stockRepo,
			StockHistoryRepository: stockHistoryRepo,
		},
	}
}
