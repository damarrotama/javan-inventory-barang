package services

import (
	"javan-inventory-barang/controller"
	"javan-inventory-barang/domain"
	"javan-inventory-barang/repository"
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
}

type Domain struct {
	ProductDomain domain.ProductDomain
}

type Repository struct {
	ProductRepository repository.ProductRepository
}

// NewService wires repositories, domain, and controllers to a shared DB pool.
func NewService() *Service {
	db, err := utils.OpenPostgres()
	if err != nil {
		log.Fatal(err)
	}

	productRepo := repository.NewProductRepository(db)
	productDomain := domain.NewProductDomain(productRepo)
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
