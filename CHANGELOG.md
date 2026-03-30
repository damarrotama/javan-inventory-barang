# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- **Stock Movement (IN/OUT)** — `POST /api/v1/stocks/movement` to add or subtract product stock with validation (stock cannot go negative).
- **Auto Stock History** — every stock movement is automatically recorded in `stock_histories` asynchronously via goroutine.
- **Stock Endpoints:**
  - `GET /api/v1/stocks` — list all current stock levels.
  - `GET /api/v1/stocks/product/:product_id` — get stock for a specific product.
  - `GET /api/v1/stocks/histories` — list all stock movement history.
  - `GET /api/v1/stocks/histories/product/:product_id` — list history for a specific product.
- Stock domain, repository, and controller layers following existing clean architecture pattern.
- `TableName()` method for `StockHistory` model to ensure correct table naming.

### Changed

- `services/service.go` — wired `StockRepository`, `StockHistoryRepository`, `StockDomain`, `StockController`, and `transaction.Manager`.
- `routes/router.go` — registered stock and stock history routes.
