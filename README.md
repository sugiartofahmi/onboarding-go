# Event Organizer Backend

A REST API backend for an event management platform, built with Go 1.25, Gin, GORM, PostgreSQL, Redis, and golang-migrate.

## Prerequisites

Before getting started, make sure the following tools are installed on your machine:

| Tool | Version | Notes |
|------|---------|-------|
| [Go](https://go.dev/dl/) | 1.25+ | Required |
| [PostgreSQL](https://www.postgresql.org/download/) | 14+ | Required |
| [Redis](https://redis.io/download/) | 7+ | Required |
| [MinIO](https://min.io/download) | Latest | Optional — for local file storage |

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

