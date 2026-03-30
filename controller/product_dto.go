package controller

// CreateProductRequest is the JSON body for POST /products.
type CreateProductRequest struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Unit        string  `json:"unit"`
}

// UpdateProductRequest is the JSON body for PUT /products/:id.
type UpdateProductRequest struct {
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Unit        string  `json:"unit"`
}
