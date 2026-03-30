
CREATE TABLE stock (
    id          BIGSERIAL PRIMARY KEY,
    product_id  BIGINT NOT NULL UNIQUE REFERENCES products (id) ON DELETE CASCADE,
    quantity    NUMERIC(18, 4) NOT NULL DEFAULT 0 CHECK (quantity >= 0),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_stock_product_id ON stock (product_id);
