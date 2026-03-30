package services

import (
	"javan-inventory-barang/controller"
	"javan-inventory-barang/domain"
	"javan-inventory-barang/repository"
	"javan-inventory-barang/transaction"
	"javan-inventory-barang/utils/database"
	"javan-inventory-barang/utils/logger"
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
	logger, err := logger.NewLogger()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.OpenPostgres()
	if err != nil {
		log.Fatal(err)
	}

	txManager := transaction.NewManager(db)

	// repositories
	productRepo := repository.NewProductRepository(logger, db)
	stockRepo := repository.NewStockRepository(db)
	stockHistoryRepo := repository.NewStockHistoryRepository(db)

	// domains
	productDomain := domain.NewProductDomain(logger, txManager, productRepo)
	stockDomain := domain.NewStockDomain(stockRepo, stockHistoryRepo, productRepo, txManager)

	// controllers
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
