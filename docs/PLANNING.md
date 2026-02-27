# event-backend — Project Planning

## Tech Stack

| Komponen | Pilihan |
|---|---|
| Language | Go 1.25 |
| Module Name | `event-backend` |
| HTTP Framework | Gin |
| ORM | GORM |
| Database | PostgreSQL |
| Cache | Redis |
| Migration | golang-migrate |
| Password Hashing | Argon2id |
| JWT | golang-jwt/v5 |
| UUID | google/uuid v7 |
| Env Loader | godotenv + os.Getenv |
| Logger | logr + stdr |
| Validation | go-playground/validator/v10 |
| Storage | S3 / MinIO / GCS |
| Mailer | SMTP / SendGrid |

---

## Project Structure

```
event-backend/
│
├── main.go
├── Dockerfile
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── README.md
│
├── app/
│   ├── auth/
│   │   ├── dtos/
│   │   │   ├── auth_register_request_dto.go
│   │   │   ├── auth_login_request_dto.go
│   │   │   └── auth_login_response_dto.go
│   │   ├── interfaces/
│   │   │   ├── auth_query_repository_interface.go
│   │   │   ├── auth_store_repository_interface.go
│   │   │   └── auth_service_interface.go
│   │   ├── repositories/
│   │   │   ├── auth_query_repository.go
│   │   │   └── auth_store_repository.go
│   │   └── services/
│   │       └── auth_service.go
│   │
│   ├── user/
│   │   ├── dtos/
│   │   │   ├── user_update_request_dto.go
│   │   │   ├── user_upgrade_organizer_request_dto.go
│   │   │   └── user_response_dto.go
│   │   ├── interfaces/
│   │   │   ├── user_query_repository_interface.go
│   │   │   ├── user_store_repository_interface.go
│   │   │   └── user_service_interface.go
│   │   ├── repositories/
│   │   │   ├── user_query_repository.go
│   │   │   └── user_store_repository.go
│   │   └── services/
│   │       └── user_service.go
│   │
│   ├── role/
│   │   ├── dtos/
│   │   │   ├── role_create_request_dto.go
│   │   │   ├── role_update_request_dto.go
│   │   │   └── role_response_dto.go
│   │   ├── interfaces/
│   │   │   ├── role_query_repository_interface.go
│   │   │   ├── role_store_repository_interface.go
│   │   │   └── role_service_interface.go
│   │   ├── repositories/
│   │   │   ├── role_query_repository.go
│   │   │   └── role_store_repository.go
│   │   └── services/
│   │       └── role_service.go
│   │
│   ├── category/
│   │   ├── dtos/
│   │   │   ├── category_create_request_dto.go
│   │   │   ├── category_update_request_dto.go
│   │   │   └── category_response_dto.go
│   │   ├── interfaces/
│   │   │   ├── category_query_repository_interface.go
│   │   │   ├── category_store_repository_interface.go
│   │   │   └── category_service_interface.go
│   │   ├── repositories/
│   │   │   ├── category_query_repository.go
│   │   │   └── category_store_repository.go
│   │   └── services/
│   │       └── category_service.go
│   │
│   ├── event/
│   │   ├── dtos/
│   │   │   ├── event_create_request_dto.go
│   │   │   ├── event_update_request_dto.go
│   │   │   ├── event_filter_request_dto.go
│   │   │   └── event_response_dto.go
│   │   ├── interfaces/
│   │   │   ├── event_query_repository_interface.go
│   │   │   ├── event_store_repository_interface.go
│   │   │   └── event_service_interface.go
│   │   ├── repositories/
│   │   │   ├── event_query_repository.go
│   │   │   └── event_store_repository.go
│   │   └── services/
│   │       └── event_service.go
│   │
│   ├── event_ticket/
│   │   ├── dtos/
│   │   │   ├── event_ticket_create_request_dto.go
│   │   │   ├── event_ticket_update_request_dto.go
│   │   │   └── event_ticket_response_dto.go
│   │   ├── interfaces/
│   │   │   ├── event_ticket_query_repository_interface.go
│   │   │   ├── event_ticket_store_repository_interface.go
│   │   │   └── event_ticket_service_interface.go
│   │   ├── repositories/
│   │   │   ├── event_ticket_query_repository.go
│   │   │   └── event_ticket_store_repository.go
│   │   └── services/
│   │       └── event_ticket_service.go
│   │
│   └── event_registration/
│       ├── dtos/
│       │   ├── event_registration_create_request_dto.go
│       │   ├── event_registration_cancel_request_dto.go
│       │   └── event_registration_response_dto.go
│       ├── interfaces/
│       │   ├── event_registration_query_repository_interface.go
│       │   ├── event_registration_store_repository_interface.go
│       │   └── event_registration_service_interface.go
│       ├── repositories/
│       │   ├── event_registration_query_repository.go
│       │   └── event_registration_store_repository.go
│       └── services/
│           └── event_registration_service.go
│
├── entities/
│   ├── user_entity.go
│   ├── role_entity.go
│   ├── category_entity.go
│   ├── event_entity.go
│   ├── event_ticket_entity.go
│   └── event_registration_entity.go
│
├── presentation/
│   └── http/
│       └── controllers/
│           ├── auth_controller.go
│           ├── user_controller.go
│           ├── role_controller.go
│           ├── category_controller.go
│           ├── event_controller.go
│           ├── event_ticket_controller.go
│           └── event_registration_controller.go
│
├── migration/
│   ├── migration.go
│   └── files/
│       ├── 20240101000001_create_roles.up.sql
│       ├── 20240101000001_create_roles.down.sql
│       ├── 20240101000002_create_users.up.sql
│       ├── 20240101000002_create_users.down.sql
│       ├── 20240101000003_create_categories.up.sql
│       ├── 20240101000003_create_categories.down.sql
│       ├── 20240101000004_create_events.up.sql
│       ├── 20240101000004_create_events.down.sql
│       ├── 20240101000005_create_event_tickets.up.sql
│       ├── 20240101000005_create_event_tickets.down.sql
│       ├── 20240101000006_create_event_registrations.up.sql
│       └── 20240101000006_create_event_registrations.down.sql
│
├── seeder/
│   ├── executor.go
│   ├── interface.go
│   ├── role_seeder.go
│   ├── category_seeder.go
│   └── user_seeder.go
│
├── constant/
│   ├── enums/
│   │   └── status_enum.go
│   ├── messages/
│   │   ├── error_message_constant.go
│   │   └── success_message_constant.go
│   ├── pagination/
│   │   └── pagination_constant.go
│   └── http/
│       └── http_constant.go
│
└── infrastructure/
    ├── config/
    │   ├── loader.go
    │   ├── app_config.go
    │   ├── database_config.go
    │   ├── jwt_config.go
    │   ├── redis_config.go
    │   ├── storage_config.go
    │   └── mailer_config.go
    │
    ├── database/
    │   └── database_connection.go
    │
    ├── redis/
    │   ├── constants/
    │   │   └── cache_key_constant.go
    │   ├── interfaces/
    │   │   └── redis_interface.go
    │   └── services/
    │       └── redis_service.go
    │
    ├── storage/
    │   ├── interfaces/
    │   │   └── storage_interface.go
    │   └── services/
    │       └── storage_service.go
    │
    ├── mailer/
    │   ├── interfaces/
    │   │   └── mailer_interface.go
    │   └── services/
    │       └── mailer_service.go
    │
    ├── jwt/
    │   └── jwt.go
    │
    ├── middleware/
    │   ├── auth_middleware.go
    │   ├── role_middleware.go
    │   ├── logger_middleware.go
    │   └── recovery_middleware.go
    │
    ├── exception/
    │   ├── handler_exception.go
    │   ├── not_found_exception.go
    │   ├── bad_request_exception.go
    │   ├── unauthorized_exception.go
    │   ├── forbidden_exception.go
    │   ├── conflict_exception.go
    │   └── server_error_exception.go
    │
    ├── response/
    │   └── response.go
    │
    ├── pagination/
    │   └── pagination.go
    │
    └── utils/
        └── slug_utils.go
```

---

## Architecture & Dependency Rule

```
presentation/controllers
  → app/{domain}/interfaces (service)
  → infrastructure/middleware
  → infrastructure/response

app/{domain}/services
  → app/{domain}/interfaces (repo)
  → entities
  → constant/*
  → infrastructure/*

app/{domain}/repositories
  → entities
  → infrastructure/database

entities
  → gorm, uuid, time only

infrastructure/*
  → infrastructure/config
  → third party only

constant/*
  → standard library only
```

---

## Database Schema

### Tables (6 tabel)

| Table | Keterangan |
|---|---|
| `roles` | Role master (attendee, organizer, admin) |
| `users` | User dengan role_id FK |
| `categories` | Kategori event, dikelola admin |
| `events` | Event milik organizer |
| `event_tickets` | Tipe tiket per event |
| `event_registrations` | Attendee yang daftar ke tiket |

### Relasi

```
roles ──< users >──────────────< event_registrations
                │                         │
                └──< events >──< event_tickets
                        │
                    categories
```

| Table | Belongs To | Has Many |
|---|---|---|
| `users` | `roles` | `events`, `event_registrations` |
| `events` | `users`, `categories` | `event_tickets` |
| `event_tickets` | `events` | `event_registrations` |
| `event_registrations` | `users`, `event_tickets` | - |
| `categories` | - | `events` |

### Field Detail

**`roles`**
```
id, name, description, created_at, updated_at, deleted_at
```

**`users`**
```
id, role_id, name, email, password, is_active, created_at, updated_at, deleted_at
```

**`categories`**
```
id, name, slug, created_at, updated_at, deleted_at
```

**`events`**
```
id, user_id, category_id, title, slug, description, location,
start_date, end_date, status, created_at, updated_at, deleted_at
```

**`event_tickets`**
```
id, event_id, type, price, quota, registered_count, created_at, updated_at, deleted_at
```

**`event_registrations`**
```
id, user_id, event_ticket_id, status, created_at, updated_at, deleted_at
UNIQUE(user_id, event_ticket_id)
```

---

## Role & RBAC System

### Roles

| Role | Deskripsi |
|---|---|
| `attendee` | User default saat register |
| `organizer` | Upgrade dari attendee via POST /users/upgrade |
| `admin` | Dibuat via seeder |

### JWT Claims

```go
type JWTClaims struct {
    UserID string `json:"user_id"`
    Role   string `json:"role"`    // "attendee" | "organizer" | "admin"
}
```

### Role Upgrade Flow

```
Register → role: attendee
POST /users/upgrade → role: organizer → return new JWT
```

### Permission Matrix

| Endpoint | Method | attendee | organizer | admin |
|---|---|---|---|---|
| /auth/register | POST | ✅ | ✅ | ✅ |
| /auth/login | POST | ✅ | ✅ | ✅ |
| /users/me | GET | ✅ | ✅ | ✅ |
| /users/me | PUT | ✅ | ✅ | ✅ |
| /users/upgrade | POST | ✅ | ❌ | ❌ |
| /users | GET | ❌ | ❌ | ✅ |
| /categories | GET | ✅ | ✅ | ✅ |
| /categories | POST | ❌ | ❌ | ✅ |
| /categories/:id | PUT | ❌ | ❌ | ✅ |
| /categories/:id | DELETE | ❌ | ❌ | ✅ |
| /events | GET | ✅ | ✅ | ✅ |
| /events | POST | ❌ | ✅ | ✅ |
| /events/:id | PUT | ❌ | ✅ (own) | ✅ |
| /events/:id | DELETE | ❌ | ✅ (own) | ✅ |
| /events/:id/publish | POST | ❌ | ✅ (own) | ✅ |
| /events/:id/tickets | GET | ✅ | ✅ | ✅ |
| /events/:id/tickets | POST | ❌ | ✅ (own) | ✅ |
| /events/:id/tickets/:id | PUT | ❌ | ✅ (own) | ✅ |
| /events/:id/tickets/:id | DELETE | ❌ | ✅ (own) | ✅ |
| /events/:id/tickets/:id/register | POST | ✅ | ❌ | ❌ |
| /registrations | GET | ✅ (own) | ✅ (incoming) | ✅ |
| /registrations/:id/cancel | POST | ✅ (own) | ❌ | ✅ |

---

## Enums

```go
// TicketType — disimpan int di DB
TicketTypeRegular = 1  → label: "Regular"
TicketTypeVIP     = 2  → label: "VIP"
TicketTypeVIPPlus = 3  → label: "VIP Plus"

// EventStatus — disimpan int di DB
EventStatusDraft     = 1  → label: "Draft"
EventStatusPublished = 2  → label: "Published"
EventStatusCancelled = 3  → label: "Cancelled"
EventStatusCompleted = 4  → label: "Completed"

// RegistrationStatus — disimpan int di DB
RegistrationStatusPending   = 1  → label: "Pending"
RegistrationStatusConfirmed = 2  → label: "Confirmed"
RegistrationStatusCancelled = 3  → label: "Cancelled"

// RoleType
RoleTypeAttendee  = 1
RoleTypeOrganizer = 2
RoleTypeAdmin     = 3
```

---

## Business Logic

| Skenario | Logic |
|---|---|
| Register event | Cek quota > registered_count, cegah double registration |
| Publish event | Cek event punya minimal 1 tiket |
| Cancel registration | Cek status masih pending/confirmed, decrement registered_count |
| Upgrade organizer | Cek role masih attendee, update role, return new JWT |
| Delete event | Cek tidak ada registrasi aktif |

---

## Race Condition Mitigations

| Skenario | Solusi |
|---|---|
| Double registration | UNIQUE(user_id, event_ticket_id) constraint |
| Oversell quota | SELECT FOR UPDATE + atomic registered_count increment |
| Duplicate upgrade | Cek role di service sebelum update |

---

## Seeder Data

**Roles:** attendee, organizer, admin

**Categories:**
```
Music, Business & Seminar, Sports & Fitness, Food & Drink,
Science & Tech, Arts & Culture, Health & Wellness,
Community, Education, Networking
```

**Users:**
```
admin@event.com → role: admin
```

---

## Naming Conventions

| Komponen | Convention | Contoh |
|---|---|---|
| Folder | `lowercase` | `repositories/` |
| File | `snake_case.go` | `event_ticket_repository.go` |
| Interface | `{Domain}{Layer}Interface` | `EventTicketServiceInterface` |
| Entity | `{Domain}Entity` | `EventTicketEntity` |
| DTO Request | `{Domain}{Action}RequestDto` | `EventTicketCreateRequestDto` |
| DTO Response | `{Domain}ResponseDto` | `EventTicketResponseDto` |
| Constructor | `New{Name}` | `NewEventTicketService` |
| Child table | `{parent}_{child}` | `event_tickets`, `event_registrations` |

---

## Dependencies

```
github.com/gin-gonic/gin
gorm.io/gorm
gorm.io/driver/postgres
github.com/golang-jwt/jwt/v5
golang.org/x/crypto
github.com/google/uuid
github.com/go-playground/validator/v10
github.com/joho/godotenv
github.com/go-logr/logr
github.com/go-logr/stdr
github.com/redis/go-redis/v9
github.com/golang-migrate/migrate/v4
github.com/golang-migrate/migrate/v4/database/postgres
github.com/golang-migrate/migrate/v4/source/file
github.com/lib/pq
```

---

## Implementation Steps

| Step | Keterangan | Status |
|---|---|---|
| Step 1 | Setup project & go modules | ⬜ Todo |
| Step 2 | Setup config & environment | ⬜ Todo |
| Step 3 | Setup database & migration | ⬜ Todo |
| Step 4 | Entities (GORM models) | ⬜ Todo |
| Step 5 | Seeder | ⬜ Todo |
| Step 6 | Infrastructure (jwt, redis, dll) | ⬜ Todo |
| Step 7 | Domain auth (register & login) | ⬜ Todo |
| Step 8 | Middleware (auth, role, logger, recovery) | ⬜ Todo |
| Step 9 | Domain user | ⬜ Todo |
| Step 10 | Domain role | ⬜ Todo |
| Step 11 | Domain category | ⬜ Todo |
| Step 12 | Domain event | ⬜ Todo |
| Step 13 | Domain event ticket | ⬜ Todo |
| Step 14 | Domain event registration | ⬜ Todo |
