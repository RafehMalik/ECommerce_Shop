# learning-go-shop

A Go-based e-commerce backend API built with Gin and GORM. This project demonstrates a simple shopping backend featuring user authentication, product/category management, cart/order operations, and file uploads.

## Features

- User registration, login, logout, and JWT-based authentication
- Refresh token support
- Admin routes for creating, updating, and deleting categories and products
- Product image upload support via local filesystem or AWS S3
- Shopping cart endpoints for adding, updating, and removing items
- Order creation and order retrieval for authenticated users
- Public product and category listing endpoints
- PostgreSQL database support with migrations
- Docker Compose setup for Postgres and LocalStack

## Tech stack

- Go 1.25
- Gin Web Framework
- GORM ORM
- PostgreSQL
- JWT authentication
- AWS SDK v2 for optional S3 uploads
- Zerolog for structured logging
- Golang Migrate for database migrations

## Project structure

- `cmd/api/main.go` - application entry point
- `internal/config` - environment configuration loader
- `internal/database` - database connection setup
- `internal/server` - route definitions and HTTP handlers
- `internal/services` - business logic services
- `internal/models` - GORM data models
- `internal/providors` - upload providers for local files and S3
- `db/migrations` - schema migrations
- `uploads/` - local upload destination
- `docker/docker-compose.yml` - local Docker services

## Requirements

- Go 1.25
- PostgreSQL (or Docker Compose)
- Docker and Docker Compose (optional)

## Environment variables

This project loads environment variables from the shell and `.env` file using `godotenv`.

Common variables:

- `PORT` - HTTP server port (default `8080`)
- `GIN_MODE` - Gin mode (`debug`, `release`, etc.)
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`, `DB_SSL_MODE`
- `JWT_SECRET`
- `JWT_EXPIRES_IN`, `REFRESH_TOKEN_EXPIRES_IN`
- `UPLOAD_PATH` - local upload folder (default `./uploads`)
- `UPLOAD_PROVIDER` - `local` or `s3`
- `AWS_REGION`, `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`, `AWS_S3_BUCKET`, `AWS_S3_ENDPOINT`

## Quick start

### Run locally

1. Create a `.env` file or set environment variables.
2. Run database migrations with:

```bash
make migrate-up
```

3. Start the application:

```bash
go run ./cmd/api
```

4. The API is available at `http://localhost:8080`.

### Run with Docker Compose

```bash
docker compose -f docker/docker-compose.yml up -d
```

This will start:

- PostgreSQL on port `5432`
- LocalStack for S3/SQS on port `4566`
- An Nginx container on port `8081`

> Note: The application itself is not started automatically by Docker Compose in this configuration.

## Makefile commands

- `make build` - Build the Go binaries
- `make run` - Run the application
- `make dev` - Run the application in development mode
- `make lint` - Run code linting
- `make format` - Format Go code
- `make migrate-up` - Apply database migrations
- `make migrate-down` - Roll back database migrations
- `make docker-up` - Start Docker Compose services
- `make docker-down` - Stop Docker Compose services

## API endpoints

Public endpoints:

- `GET /health`
- `GET /api/v1/categories`
- `GET /api/v1/products`
- `GET /api/v1/products/:id`

Authentication endpoints:

- `POST /api/v1/auth/register`
- `POST /api/v1/auth/login`
- `POST /api/v1/auth/logout`
- `POST /api/v1/auth/refresh`

Protected endpoints (require JWT):

- `GET /api/v1/user/profile`
- `PUT /api/v1/user/profile`
- `GET /api/v1/cart`
- `POST /api/v1/cart/items`
- `PUT /api/v1/cart/items/:id`
- `DELETE /api/v1/cart/items/:id`
- `GET /api/v1/orders`
- `GET /api/v1/orders/:id`
- `POST /api/v1/orders`

Admin-only endpoints:

- `POST /api/v1/categories`
- `PUT /api/v1/categories/:id`
- `DELETE /api/v1/categories/:id`
- `POST /api/v1/products`
- `PUT /api/v1/products/:id`
- `DELETE /api/v1/products/:id`
- `POST /api/v1/products/:id/images`

## Notes

- Local uploads are saved to the configured `UPLOAD_PATH`.
- S3 uploads can be configured using AWS credentials and endpoint.
- The repository uses database migrations in `db/migrations`.

## License

This project is provided for learning purposes.
