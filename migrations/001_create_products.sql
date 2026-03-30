-- Products: master data for items you track in inventory.
-- Adjust columns when you define the final product shape (SKU, category, etc.).

CREATE TABLE products (
    id          BIGSERIAL PRIMARY KEY,
    sku         VARCHAR(64)  NOT NULL UNIQUE,
    name        VARCHAR(255) NOT NULL,
    description TEXT,
    unit        VARCHAR(32)  NOT NULL DEFAULT 'pcs',
    created_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_products_name ON products (name);
