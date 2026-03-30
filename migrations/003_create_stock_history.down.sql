BEGIN;

DROP INDEX IF EXISTS idx_stock_history_created_at;
DROP INDEX IF EXISTS idx_stock_history_product_id;

DROP TABLE IF EXISTS stock_history;

DROP TYPE IF EXISTS stock_movement_type;

COMMIT;