# Javan Inventory Barang

REST API for managing **inventory items** (“barang”): product master data backed by PostgreSQL. Built with Go, [Fiber](https://github.com/gofiber/fiber), and [GORM](https://gorm.io/) in a layered layout (repository → domain → controller).

## Features

- CRUD HTTP API for products (SKU, name, description, unit).
- OpenAPI/Swagger UI for interactive documentation.
- PostgreSQL with optional GORM auto-migration for `products`, `stocks`, and `stock_histories` (see [Database](#database)).
- Configuration via `.env` and environment variables ([Viper](https://github.com/spf13/viper)).

## Requirements

- Go **1.25.3** or compatible (see `go.mod`).
- PostgreSQL **14+** (or any version supported by `pgx`/`gorm` drivers).

## Quick start

1. **Clone** the repository and enter the project directory.

2. **Create a database** in PostgreSQL for this application.

3. **Configure environment** — copy the example file and adjust values:

   ```bash
   cp .env.example .env
   ```

   Set at least `PORT`, `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, and `DB_NAME`.

4. **Schema** — choose one:

   - **Development (recommended):** set `DB_AUTO_MIGRATE=true` in `.env` so GORM creates/updates tables on startup from the models in `migrations/migration.go`.
   - **Manual:** run the SQL files in `migrations/` in order (`001_…`, `002_…`, `003_…`) if you prefer fixed migrations without auto-migrate.

   Optional: `DB_TABLE_PREFIX` — table name prefix used by GORM’s naming strategy (empty if unset).

5. **Run the server:**

   ```bash
   go run .
   ```

   The API listens on the port from `PORT` (default in `.env.example`: **8000**).

## API

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/` | Simple health-style response |
| `GET` | `/api/v1/products` | List products |
| `POST` | `/api/v1/products` | Create product |
| `GET` | `/api/v1/products/:id` | Get product by ID |
| `PUT` | `/api/v1/products/:id` | Update product |
| `DELETE` | `/api/v1/products/:id` | Delete product |

**Swagger UI:** after starting the server, open:

`http://localhost:<PORT>/swagger/index.html`

(Replace `<PORT>` with your configured `PORT`.)

JSON bodies for create/update use `sku`, `name`, optional `description`, and `unit` (see `controller/product_dto.go`).

## Project layout

| Path | Role |
|------|------|
| `main.go` | Fiber app, config init, listen |
| `routes/router.go` | Routes and Swagger mount |
| `controller/` | HTTP handlers and request DTOs |
| `domain/` | Business logic |
| `repository/` | Data access |
| `model/` | GORM models |
| `utils/` | Config (`env`) and PostgreSQL (`database`) |
| `migrations/` | GORM model list for auto-migrate + ordered SQL scripts |
| `docs/` | Generated Swagger (`swagger.json`, `swagger.yaml`, `docs.go`) |

## Regenerating Swagger docs

If you change handler annotations (`@Summary`, `@Router`, etc.), regenerate with [swag](https://github.com/swaggo/swag):

```bash
go install github.com/swaggo/swag/cmd/swag@latest
swag init
```

This updates the files under `docs/`.

## License

Specify your license here if applicable.
