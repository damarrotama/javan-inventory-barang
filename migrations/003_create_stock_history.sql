-- Stock history: append-only log of every quantity change (receipts, sales, adjustments).

CREATE TYPE stock_movement_type AS ENUM ('IN', 'OUT', 'ADJUSTMENT');

CREATE TABLE stock_history (
    id              BIGSERIAL PRIMARY KEY,
    product_id      BIGINT NOT NULL REFERENCES products (id) ON DELETE CASCADE,
    stock_id        BIGINT REFERENCES stock (id) ON DELETE SET NULL,
    movement_type   stock_movement_type NOT NULL,
    quantity_delta  NUMERIC(18, 4) NOT NULL,
    quantity_after  NUMERIC(18, 4),
    reference       VARCHAR(128),
    note            TEXT,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_stock_history_product_id ON stock_history (product_id);
CREATE INDEX idx_stock_history_created_at ON stock_history (created_at);
