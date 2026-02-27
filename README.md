# Event Organizer Backend

A REST API backend for an event management platform, built with Go 1.25, Gin, GORM, PostgreSQL, Redis, and golang-migrate.

## Prerequisites

Before getting started, make sure the following tools are installed on your machine:

| Tool | Version | Notes |
|------|---------|-------|
| [Go](https://go.dev/dl/) | 1.25+ | Required for local setup |
| [PostgreSQL](https://www.postgresql.org/download/) | 14+ | Required for local setup (provided by Docker) |
| [Redis](https://redis.io/download/) | 7+ | Required for local setup (provided by Docker) |
| [MinIO](https://min.io/download) | Latest | Optional for local setup (provided by Docker) |
| [Docker](https://docs.docker.com/get-docker/) | 20.10+ | Required for Docker setup |
| [Docker Compose](https://docs.docker.com/compose/install/) | 2.0+ | Required for Docker setup (included with Docker Desktop) |

> **Note:** If you use [Docker Setup](#docker-setup), you do **not** need to install Go, PostgreSQL, Redis, or MinIO on your machine.

---

## Local Setup

### Step 1: Clone the repository

```bash
git clone <repository-url>
cd event-backend## Prerequisites
```

### Step 2: Configure environment variables

```bash
cp .env-example .env
```

Open `.env` and fill in the required values. See the [Environment Variables Reference](#environment-variables-reference) section below for details.

### Step 3: Install dependencies

```bash
go mod download
```

### Step 4: Create the PostgreSQL database

Connect to your PostgreSQL instance and create the database:

```sql
CREATE DATABASE event_backend;
```

### Step 5: Run database migrations

```bash
go run main.go --migration=true --exec=up
```

This applies all pending migrations and creates all required tables.

### Step 6: Start the server

```bash
go run main.go
```

The server will start on the port defined by `APP_PORT` (default: `8080`).

---

## Docker Setup

If you prefer running the application and its dependencies via Docker, follow these steps instead of [Local Setup](#local-setup).

### Compose Files

| File | Environment | Env File | Description |
|------|-------------|----------|-------------|
| `docker-compose.yml` | Development | `.env` | All ports exposed to host, hot-reload friendly |
| `docker-compose.staging.yml` | Staging | `.env.staging` | Internal service ports hidden, `restart: always` |
| `docker-compose.production.yml` | Production | `.env.production` | Resource limits, logging config, hardened security |

Each compose file reads its environment variables from a **separate env file**. The `--env-file` flag tells Docker Compose which file to use for `${...}` variable substitution in the compose file, and the `env_file` directive inside the compose file injects the same variables into the container.

### Services Overview

| Service | Port (Dev) | Port (Staging) | Port (Production) | Description |
|---------|------------|----------------|--------------------|-------------|
| `app` | `8080` | `8081` | `8082` | Go application |
| `postgres` | `5432` | not exposed | not exposed | PostgreSQL 17 |
| `redis` | `6379` | not exposed | not exposed | Redis 7 |
| `minio` | `9000`, `9001` | `9002` (console only) | not exposed | S3-compatible object storage |
| `minio-setup` | — | — | — | One-time bucket creation |

> **Note:** Ports are different per environment so all three can run simultaneously on the same machine without conflicts.

### Step 1: Configure environment variables

Copy the appropriate example file for your environment:

**Development:**

```bash
cp .env-example .env
```

**Staging:**

```bash
cp .env.staging.example .env.staging
```

**Production:**

```bash
cp .env.production.example .env.production
```

All example files are pre-configured for Docker (hosts point to container service names: `postgres`, `redis`, `minio`). Adjust credentials as needed.

> **Warning:** For staging and production, make sure to set strong values for `DB_PASSWORD`, `JWT_SECRET`, `AWS_S3_ACCESS_KEY`, and `AWS_S3_SECRET_KEY`. The production example leaves these **empty** — they must be filled in before starting.

### Step 2: Start all services

**Development:**

```bash
docker compose -f docker-compose.yml up -d
```

**Staging:**

```bash
docker compose --env-file .env.staging -f docker-compose.staging.yml up -d
```

**Production:**

```bash
docker compose --env-file .env.production -f docker-compose.production.yml up -d
```

To rebuild the application image (after code changes), add `--build`:

```bash
docker compose -f docker-compose.yml up -d --build
```

> **Note:** Development does not need `--env-file` because Docker Compose automatically reads `.env` from the project root.

### Step 3: Run database migrations

**Development:**

```bash
docker compose -f docker-compose.yml exec app ./event-backend --migration=true --exec=up
```

**Staging:**

```bash
docker compose --env-file .env.staging -f docker-compose.staging.yml exec app ./event-backend --migration=true --exec=up
```

**Production:**

```bash
docker compose --env-file .env.production -f docker-compose.production.yml exec app ./event-backend --migration=true --exec=up
```

### Step 4: Run database seeders (optional)

```bash
docker compose -f docker-compose.yml exec app ./event-backend --dbseed=true
```

### Step 5: Verify

Check application logs:

```bash
docker compose -f docker-compose.yml logs -f app
```

| Environment | API | MinIO Console |
|-------------|-----|---------------|
| Development | [http://localhost:8080](http://localhost:8080) | [http://localhost:9001](http://localhost:9001) |
| Staging | [http://localhost:8081](http://localhost:8081) | [http://localhost:9002](http://localhost:9002) |
| Production | [http://localhost:8082](http://localhost:8082) | not exposed |

MinIO Console credentials (dev/staging): `minioadmin` / `minioadmin`

### Common Docker Commands

The table below shows **development** commands. For staging or production, add `--env-file` and swap the compose file:

```bash
# Staging pattern
docker compose --env-file .env.staging -f docker-compose.staging.yml <command>

# Production pattern
docker compose --env-file .env.production -f docker-compose.production.yml <command>
```

| Command | Description |
|---------|-------------|
| `docker compose -f docker-compose.yml up -d` | Start all services in background |
| `docker compose -f docker-compose.yml up -d --build` | Rebuild image and start all services |
| `docker compose -f docker-compose.yml down` | Stop all services |
| `docker compose -f docker-compose.yml down -v` | Stop all services and remove volumes (reset data) |
| `docker compose -f docker-compose.yml ps` | Show status of all services |
| `docker compose -f docker-compose.yml logs -f app` | Follow application logs |
| `docker compose -f docker-compose.yml logs -f postgres` | Follow database logs |
| `docker compose -f docker-compose.yml restart app` | Restart only the application |
| `docker compose -f docker-compose.yml exec app ./event-backend --migration=true --exec=up` | Run migrations |
| `docker compose -f docker-compose.yml exec app ./event-backend --migration=true --exec=down` | Rollback last migration |
| `docker compose -f docker-compose.yml exec app ./event-backend --migration=true --exec=fresh` | Fresh migration (drop all + re-apply) |
| `docker compose -f docker-compose.yml exec app ./event-backend --dbseed=true` | Run all seeders |
| `docker compose -f docker-compose.yml exec app ./event-backend --dbseed=true --class=RoleSeeder` | Run specific seeder |

### Environment Files

| File | Example File | Git Tracked | Description |
|------|-------------|-------------|-------------|
| `.env` | `.env-example` | No | Development environment variables |
| `.env.staging` | `.env.staging.example` | No | Staging environment variables |
| `.env.production` | `.env.production.example` | No | Production environment variables |

### Environment Differences

| Feature | Development | Staging | Production |
|---------|-------------|---------|------------|
| `APP_ENV` | `development` | `release` | `release` |
| Restart policy | `unless-stopped` | `always` | `always` |
| Internal ports exposed | Yes | No | No |
| Resource limits | No | No | Yes |
| Log rotation | No | No | Yes (`10m`, 3 files) |
| Redis `maxmemory` | Default | Default | `200mb` (LRU eviction) |
| MinIO public access | Bucket set to download | Bucket set to download | No public access |
| DB pool (idle/open) | 5 / 10 | 5 / 10 | 10 / 25 |
| `DB_PASSWORD` | `test` | `test` | **Required** (empty) |
| `JWT_SECRET` | `change-me-...` | `change-me-...` | **Required** (empty) |
| `AWS_S3_ACCESS_KEY` | `minioadmin` | `minioadmin` | **Required** (empty) |

---

## Migration System

This project uses [golang-migrate/v4](https://github.com/golang-migrate/migrate) with file-based SQL migrations.

### How It Works

- Migration files are stored in `migration/files/` as plain SQL files
- Each migration has two files: `.up.sql` (apply) and `.down.sql` (rollback)
- Migration state is tracked in the `schema_migrations` table in PostgreSQL
- The `uuid-ossp` PostgreSQL extension is automatically created before migrations run
- Dirty migration states (from interrupted migrations) are auto-recovered before each run

### Migration Commands

All migration commands use the `--migration=true` flag:

| Command | Description |
|---------|-------------|
| `go run main.go --migration=true --exec=up` | Apply all pending migrations |
| `go run main.go --migration=true --exec=down` | Roll back the last applied migration |
| `go run main.go --migration=true --exec=fresh` | Drop all tables and re-apply all migrations from scratch |
| `go run main.go --migration=true --exec=create --fileName=<name>` | Create a new empty migration file pair |

### Creating a New Migration

To scaffold a new migration, run the `create` command — **no database connection is required**:

```bash
go run main.go --migration=true --exec=create --fileName=create_table_products
```

This creates two empty files with an auto-incremented sequence number:

```
migration/files/000007_create_table_products.up.sql
migration/files/000007_create_table_products.down.sql
```

Fill in the `.up.sql` with your `CREATE TABLE` or `ALTER TABLE` statements, and the `.down.sql` with the corresponding rollback.

### Migration File Naming Convention

```
{6-digit-sequence}_{verb}_{subject}
```

Examples:
- `000007_create_table_products`
- `000008_add_column_users_phone`
- `000009_drop_index_events_status`

### Existing Migrations

| # | File | Description |
|---|------|-------------|
| 1 | `000001_create_table_roles` | Roles for RBAC (attendee, organizer, admin) |
| 2 | `000002_create_table_users` | User accounts with role association |
| 3 | `000003_create_table_categories` | Event categories |
| 4 | `000004_create_table_events` | Events with organizer, category, and status |
| 5 | `000005_create_table_event_tickets` | Ticket types per event with quota tracking |
| 6 | `000006_create_table_event_registrations` | Registration records linking users to tickets |

---

## Database Seeder

Seeders populate reference/static data (roles, categories, users) needed for development.

### Seeder Commands

| Command | Description |
|---------|-------------|
| `go run main.go --dbseed=true` | Run all seeders (truncates existing data first) |
| `go run main.go --dbseed=true --class=RoleSeeder` | Run a single seeder by name |
| `go run main.go --dbseed=true --class=RoleSeeder,CategorySeeder` | Run multiple specific seeders |

> **Warning:** Running without `--class` deletes all rows in `users`, `categories`, and `roles` before re-seeding.

### Available Seeders

| Name | Description |
|------|-------------|
| `RoleSeeder` | Inserts Admin, Organizer, Attendee roles |
| `CategorySeeder` | Inserts 8 default event categories |
| `UserSeeder` | Inserts one user per role with hashed passwords |

Seed data files live in `seeder/files/*.json`.

---

## Running the Application

**Development (run directly):**

```bash
go run main.go
```

**Build a binary:**

```bash
go build -o event-backend .
```

**Run the binary:**

```bash
./event-backend
```

**Build and run (production-like):**

```bash
go build -o event-backend . && ./event-backend
```

The server supports graceful shutdown — it waits up to 5 seconds for in-flight requests to complete when it receives `SIGINT` or `SIGTERM`.

---

## Running Tests

Run the full test suite:

```bash
go test ./...
```

Run only domain-layer tests:

```bash
go test ./app/...
```

Run with verbose output (shows each test name and result):

```bash
go test -v ./app/...
```

Run with coverage report:

```bash
go test -cover ./app/...
```

Generate an HTML coverage report:

```bash
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out
```

---

## Environment Variables Reference

### Application

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `APP_NAME` | No | `event-backend` | Application name (used in logs) |
| `APP_ENV` | No | `development` | Gin mode: `development` or `release` |
| `APP_PORT` | No | `8080` | HTTP port the server listens on |

### Database

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DB_HOST` | Yes | `localhost` | PostgreSQL host |
| `DB_PORT` | Yes | `5432` | PostgreSQL port |
| `DB_USER` | Yes | `postgres` | PostgreSQL username |
| `DB_PASSWORD` | **Yes** | — | PostgreSQL password |
| `DB_NAME` | Yes | `event_backend` | Database name |
| `DB_TIMEZONE` | No | `Asia/Jakarta` | Database timezone |

### Database Connection Pool

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `DB_MAX_IDLE_CONNS` | No | `5` | Maximum idle connections |
| `DB_MAX_OPEN_CONNS` | No | `10` | Maximum open connections |
| `DB_MAX_IDLE_CONNS_IN_MINUTES` | No | `10` | Idle connection timeout (minutes) |
| `DB_MAX_LIFETIME_CONNS_IN_MINUTES` | No | `60` | Maximum connection lifetime (minutes) |

### JWT Authentication

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `JWT_SECRET` | **Yes** | — | Secret key for signing JWT tokens — **must be set manually** |
| `JWT_EXPIRED_IN` | No | `24h` | Token expiry duration (e.g. `1h`, `24h`, `7d`) |

### Redis

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `REDIS_HOST` | Yes | `localhost` | Redis host |
| `REDIS_PORT` | Yes | `6379` | Redis port |
| `REDIS_PASS` | No | — | Redis password (leave empty if no auth) |

### Storage (MinIO / AWS S3)

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `STORAGE` | No | `minio` | Storage backend (`minio` or `s3`) |
| `FILE_UPLOAD_MAX_SIZE` | No | `10` | Maximum upload size in MB |
| `FILE_UPLOAD_ALLOWED_TYPES` | No | `image/jpeg,image/png` | Comma-separated list of allowed MIME types |
| `AWS_S3_ENDPOINT` | No | — | S3/MinIO endpoint URL |
| `AWS_S3_PORT` | No | — | S3/MinIO port (for local MinIO) |
| `AWS_S3_BUCKET_NAME` | No | `event-backend` | S3/MinIO bucket name |
| `AWS_S3_ACCESS_KEY` | No | — | S3/MinIO access key |
| `AWS_S3_SECRET_KEY` | No | — | S3/MinIO secret key |
| `AWS_S3_REGION` | No | — | S3 region (e.g. `us-east-1`) |

---

## Troubleshooting

### Database connection refused

```
failed to connect to database: ...
```

- Ensure PostgreSQL is running: `pg_isready` or `systemctl status postgresql`
- Verify `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, and `DB_NAME` in your `.env` file
- Confirm the database exists: `psql -U postgres -l`

### Redis connection failed

```
failed to connect to redis: ...
```

- Ensure Redis is running: `redis-cli ping` (should return `PONG`)
- Verify `REDIS_HOST`, `REDIS_PORT`, and `REDIS_PASS` in your `.env` file

### Dirty migration state

```
migration is dirty at version N, forcing...
```

This happens when a migration was interrupted partway through. The migration runner **automatically detects and recovers** from dirty states by forcing the version before re-running. No manual intervention is required.

If you want to recover manually, you can also run a fresh migration:

```bash
go run main.go --migration=true --exec=fresh
```

> **Warning:** `--exec=fresh` drops all tables and data. Only use this in development.

