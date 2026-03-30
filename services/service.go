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
}

type Domain struct {
	ProductDomain domain.ProductDomain
}

type Repository struct {
	ProductRepository repository.ProductRepository
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

	// domains
	productDomain := domain.NewProductDomain(logger, txManager, productRepo)

	// controllers
	productController := controller.NewProductController(productDomain)

	return &Service{
		Controller: Controller{
			ProductController: productController,
		},
		Domain: Domain{
			ProductDomain: productDomain,
		},
		Repository: Repository{
			ProductRepository: productRepo,
		},
	}
}
