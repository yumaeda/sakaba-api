# AGENTS.md — sakaba-api Development Guide

## Project Overview

`sakaba-api` is a Go 1.19 REST API backend for sakaba.link — a restaurant discovery platform. It is deployed as a containerised service on AWS ECS, backed by TiDB (MySQL-compatible).

Full documentation: see [`CLAUDE.md`](./CLAUDE.md).

## Quick Reference

| Layer | Package | Purpose |
|-------|---------|---------|
| Model | `src/model/` | Domain entities & DTOs (GORM structs + view types) |
| Repository | `src/repository/` | Data access — raw SQL via GORM, binary UUID conversion |
| Controller | `src/controller/` | HTTP handlers — thin wrappers returning `{"statusCode", "body"}` |
| Middleware | `src/middleware/` | JWT authentication (gin-jwt) |
| Infrastructure | `src/infrastructure/` | DB connection, S3 client, distance calculation SQL |

## Architecture Patterns

### 1. Dependency Injection via Struct Fields

Controllers embed their repositories; repositories embed the GORM DB handle.

```go
type RestaurantController struct {
    Repository repository.RestaurantRepository
}

type RestaurantRepository struct {
    TiDB *gorm.DB
}
```

Controllers are wired in `main.go` with explicit instantiation.

### 2. Binary UUIDs

TiDB stores UUIDs as binary (16 bytes) for performance. All reads use `BIN_TO_UUID(col, 1)`; all writes use `infrastructure.UUIDToBin()`. Never handle raw UUIDs at the DB layer — always convert.

### 3. Response Shape

Every controller returns:

```json
{"statusCode": 200, "body": <data>}
```

Error responses omit `body` and include an `error` field.

### 4. Distance Queries

Endpoints that accept `latitude` and `longitude` use `infrastructure.GetDistanceSQL()` — the Haversine formula (Earth radius 6371.009 km). Results are ordered `ASC` by distance.

### 5. Open-Status Check

`is_open` is computed server-side using the restaurant's `business_day_info` JSON column and the current JST time. `DAYOFWEEK(CURDATE())` selects the day slot.

## API Surface

| Method | Path | Auth | Description |
|--------|------|------|-------------|
| POST | `/login` | — | JWT login |
| GET | `/` | — | Health check |
| GET | `/health/` | — | Health check (duplicate) |
| GET | `/areas/` | — | List all areas |
| GET | `/categories/:id` | — | Categories by restaurant ID |
| GET | `/dishes/` | — | List all dishes |
| GET | `/dishes/:id` | — | Dish by ID |
| GET | `/drinks/` | — | List all drinks |
| GET | `/drinks/:id` | — | Drink by ID |
| GET | `/genres/` | — | List all genres |
| GET | `/genres/:id` | — | Genre by ID |
| GET | `/photos/:id` | — | Photos by restaurant ID |
| GET | `/latest-photos/` | — | Latest photos across all restaurants |
| GET | `/menus/:id` | — | Menus by restaurant ID |
| GET | `/restaurants/` | — | List all open restaurants |
| GET | `/restaurants/areas/:id/:latitude/:longitude` | — | Restaurants by area + distance |
| GET | `/restaurants/dishes/:id/:latitude/:longitude` | — | Restaurants by dish + distance |
| GET | `/restaurants/drinks/:id/:latitude/:longitude` | — | Restaurants by drink + distance |
| GET | `/restaurants/genres/:id/:latitude/:longitude` | — | Restaurants by genre + distance |
| GET | `/restaurant-counts/:latitude/:longitude` | — | Open restaurant counts by area |
| GET | `/rankings/` | — | List all rankings |
| GET | `/videos/` | — | Videos |
| GET | `/videos/:id` | — | Videos by restaurant ID |
| GET | `/auth/refresh_token` | Bearer | Refresh JWT |
| GET | `/auth/home` | Bearer | Admin home |
| POST | `/auth/photo/` | Bearer | Add photo |
| POST | `/auth/restaurant/` | Bearer | Add restaurant |
| POST | `/auth/restaurant-drink/` | Bearer | Link drink to restaurant |
| POST | `/auth/restaurant-genre/` | Bearer | Link genre to restaurant |
| POST | `/auth/menu/` | Bearer | Add menu |
| PUT | `/auth/menu/` | Bearer | Update menu |
| DELETE | `/auth/menu/` | Bearer | Delete menu |

## Build & Deploy

```sh
# Build Docker image
docker build --no-cache -t sakaba/api .

# Run locally
docker run --rm -d \
    -e APP_CONFIG_JSON='{"db.password":"...","db.host":"...","db.name":"...","db.user":"...","aws.s3.id":"...","aws.s3.secret":"...","aws.s3.region":"..."}' \
    -p 8080:8080 \
    sakaba/api

# Publish to ECR (requires AWS CLI)
./publish.sh
```

Environment variables:

| Variable | Purpose |
|----------|---------|
| `APP_CONFIG_JSON` | DB + S3 credentials (JSON string) |
| `TIDB_CONFIG_JSON` | Alternative DB config source (parsed in `infrastructure.ConnectToDB`) |
| `GIN_MODE` | Set to `release` in Docker |

## Linting & Quality

```sh
golint ./...
go vet ./...
shadow ./...
go mod tidy -v
```

## Adding a New Resource

Follow the existing three-layer pattern:

1. **Model** (`src/model/`) — Define the GORM struct (with `gorm:"..."` tags) and any view DTOs.
2. **Repository** (`src/repository/`) — Implement data access using raw SQL via `c.TiDB.Raw(...).Scan(&result)`. Use `infrastructure.UUIDToBin()` for inserts; `BIN_TO_UUID(..., 1)` for selects.
3. **Controller** (`src/controller/`) — Thin handler: bind JSON, call repository, return `{"statusCode", "body"}`.
4. **main.go** — Wire the controller with its repository, register routes.

## Security Notes

- Password hashing is plaintext comparison (see `AuthMiddleware.go:44` — "TODO: Use Argon2"). Do not add new plaintext-password paths.
- JWT secret key is hardcoded (`"testKey"` in `main.go`) — should be externalised.
- Distance SQL uses string interpolation — validate latitude/longitude inputs before passing to `GetDistanceSQL`.
