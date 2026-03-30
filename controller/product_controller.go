package controller

import (
	"errors"
	"javan-inventory-barang/domain"
	"javan-inventory-barang/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(products)
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
	id, err := parseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	p, err := c.productDomain.GetProductByID(ctx.UserContext(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(p)
}

// CreateProduct godoc
// @Summary Create product
// @Tags products
// @Accept json
// @Produce json
// @Param body body CreateProductRequest true "Payload"
// @Success 201 {object} model.Product
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
func (c *productController) CreateProduct(ctx *fiber.Ctx) error {
	var req CreateProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}
	if strings.TrimSpace(req.SKU) == "" || strings.TrimSpace(req.Name) == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "sku and name are required"})
	}
	unit := strings.TrimSpace(req.Unit)
	if unit == "" {
		unit = "pcs"
	}
	p := &model.Product{
		SKU:         strings.TrimSpace(req.SKU),
		Name:        strings.TrimSpace(req.Name),
		Description: req.Description,
		Unit:        unit,
	}
	if err := c.productDomain.CreateProduct(ctx.UserContext(), p); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(p)
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
	id, err := parseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req UpdateProductRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid json"})
	}
	if strings.TrimSpace(req.SKU) == "" || strings.TrimSpace(req.Name) == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "sku and name are required"})
	}
	unit := strings.TrimSpace(req.Unit)
	if unit == "" {
		unit = "pcs"
	}
	p := &model.Product{
		Base:        model.Base{ID: id},
		SKU:         strings.TrimSpace(req.SKU),
		Name:        strings.TrimSpace(req.Name),
		Description: req.Description,
		Unit:        unit,
	}
	if err := c.productDomain.UpdateProduct(ctx.UserContext(), p); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(p)
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
	id, err := parseIDParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := c.productDomain.DeleteProduct(ctx.UserContext(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func parseIDParam(ctx *fiber.Ctx) (*uuid.UUID, error) {
	raw := ctx.Params("id")
	id, err := uuid.Parse(raw)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
