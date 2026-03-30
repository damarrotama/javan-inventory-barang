package controller

import (
	"javan-inventory-barang/domain"
	"javan-inventory-barang/lib"
	"javan-inventory-barang/model"
	"javan-inventory-barang/utils/resp"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	GetProducts(ctx *fiber.Ctx) error
	GetProductByID(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProduct(ctx *fiber.Ctx) error
	DeleteProduct(ctx *fiber.Ctx) error
}

// ProductController handles HTTP for products.
type productController struct {
	productDomain domain.ProductDomain
}

// NewProductController creates a ProductController.
func NewProductController(productDomain domain.ProductDomain) ProductController {
	return &productController{productDomain: productDomain}
}

// GetProducts godoc
// @Summary List all products
// @Tags products
// @Produce json
// @Success 200 {array} model.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (c *productController) GetProducts(ctx *fiber.Ctx) error {
	products, err := c.productDomain.GetProducts(ctx.UserContext())
	if err != nil {
		return resp.ErrorHandler(ctx, err)
	}
	return resp.OK(ctx, products)
}

// GetProductByID godoc
// @Summary Get product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [get]
func (c *productController) GetProductByID(ctx *fiber.Ctx) error {
	product, err := c.productDomain.GetProductByID(ctx.UserContext(), lib.ParamsUUID(ctx))
	if err != nil {
		return resp.ErrorHandler(ctx, err)
	}
	return resp.OK(ctx, product)
}

// CreateProduct godoc
// @Summary Create product
// @Tags products
// @Accept json
// @Produce json
// @Param body body model.ProductAPI true "Payload"
// @Success 201 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (c *productController) CreateProduct(ctx *fiber.Ctx) error {
	api := new(model.ProductAPI)
	if err := ctx.BodyParser(api); err != nil {
		return resp.ErrorHandler(ctx, resp.ErrorBadRequest("invalid json"))
	}

	product, err := c.productDomain.CreateProduct(ctx.UserContext(), api)
	if err != nil {
		return resp.ErrorHandler(ctx, err)
	}
	return resp.Created(ctx, product)
}

// UpdateProduct godoc
// @Summary Update product
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID" format(uuid)
// @Param body body UpdateProductRequest true "Payload"
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [put]
func (c *productController) UpdateProduct(ctx *fiber.Ctx) error {
	api := new(model.ProductAPI)
	if err := ctx.BodyParser(api); err != nil {
		return resp.ErrorHandler(ctx, resp.ErrorBadRequest("invalid json"))
	}

	product, err := c.productDomain.UpdateProduct(ctx.UserContext(), api, lib.ParamsUUID(ctx))
	if err != nil {
		return resp.ErrorHandler(ctx, err)
	}
	return resp.OK(ctx, product)
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags products
// @Param id path string true "Product ID" format(uuid)
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [delete]
func (c *productController) DeleteProduct(ctx *fiber.Ctx) error {
	if err := c.productDomain.DeleteProduct(ctx.UserContext(), lib.ParamsUUID(ctx)); err != nil {
		return resp.ErrorHandler(ctx, err)
	}
	return resp.OK(ctx)
}
