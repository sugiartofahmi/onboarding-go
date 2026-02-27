# Coding Standards

Internal coding standards for the event-backend project. All contributors must follow these rules.

---

## Table of Contents

1. [Directory Structure](#1-directory-structure)
2. [Naming Conventions](#2-naming-conventions)
3. [Avoid Magic Values](#3-avoid-magic-values)
4. [Extract Complex Conditions](#4-extract-complex-conditions)
5. [Repository Query Design](#5-repository-query-design)
6. [Receiving Request Data in Gin](#6-receiving-request-data-in-gin)
7. [Pointer Usage Best Practices](#7-pointer-usage-best-practices)
8. [Existence Checks](#8-existence-checks)
9. [Avoid N+1 Queries](#9-avoid-n1-queries)
10. [Avoid Monster Functions](#10-avoid-monster-functions)

---

## 1. Directory Structure

Each domain lives under `app/{domain}/` and follows this layout:

```
app/{domain}/
  dtos/           # request & response DTOs
  interfaces/     # repository & service interfaces
  repositories/   # query_repository + store_repository
  services/       # business logic
```

Top-level directories:

```
entities/                         # GORM entity structs (shared across domains)
presentation/http/controllers/    # HTTP controllers per domain
constant/                         # app-wide constants and enums
infrastructure/                   # shared infra: config, DB, Redis, enums, DTOs, utils
migration/                        # migration runner and SQL files
seeder/                           # database seeders
```

### File Naming Convention

Pattern: `{domain}_{layer}_{type}.go`

| Example | Description |
|---------|-------------|
| `user_query_repository.go` | Read-only repository for users |
| `user_store_repository.go` | Write-only repository for users |
| `auth_service.go` | Service for auth domain |
| `event_response_dto.go` | Response DTO for event domain |
| `event_query_repository_interface.go` | Interface for event query repo |

Package names must be short, lowercase, no underscores: `services`, `repositories`, `dtos`.

---

## 2. Naming Conventions

### Variables & Booleans

Boolean variables must use a semantic prefix: `is`, `has`, `can`, `should`, `was`, `will`.

```go
// Bad
active := user.Status

// Good
isActive := user.Status == "active"
hasExpired := ticket.ExpiresAt.Before(time.Now())
canPurchase := user.Balance >= event.Price
```

### Functions

Use verb-based names that communicate intent:

| Verb group | Use for |
|------------|---------|
| `get` / `fetch` / `find` | Read operations |
| `create` / `generate` | Write / creation |
| `update` / `modify` | Mutation |
| `delete` / `remove` | Destruction |
| `validate` / `check` | Verification |

```go
// Bad
func handleUser() {}

// Good
func deactivateUser(id uuid.UUID) error {}
func archiveExpiredEvent(eventID uuid.UUID) error {}
func validateTicketQuota(eventID uuid.UUID, requested int) error {}
```

### Constants

Use `PascalCase` for exported constants. Group related constants in dedicated files.

```
constant/enums/status_enum.go
constant/messages/error_message_constant.go
```

```go
// constant/enums/role_enum.go
const (
    RoleAdmin    = 1
    RoleOrganizer = 2
    RoleAttendee  = 3
)

// constant/order_constant.go
const MaxItemsPerOrder = 50
```

### Files & Packages

- Files: `snake_case` — `event_query_repository.go`
- Packages: short, lowercase, no underscores — `services`, `repositories`, `dtos`

---

## 3. Avoid Magic Values

Replace raw literals with named constants.

```go
// Bad
if user.RoleID == 3 {
    // ...
}
if len(items) > 50 {
    // ...
}

// Good
const RoleAdmin = 3
const MaxItemsPerOrder = 50

if user.RoleID == RoleAdmin {
    // ...
}
if len(items) > MaxItemsPerOrder {
    // ...
}
```

This applies to strings, integers, durations, and any literal used more than once or that carries business meaning.

---

## 4. Extract Complex Conditions

Extract multi-part boolean expressions into named variables.

```go
// Bad
if user.Age >= 17 && user.KycStatus == "verified" && !user.IsBanned {
    // allow purchase
}

// Good
meetsAgeRequirement := user.Age >= MinAge
isIdentityVerified   := user.KycStatus == "verified"
isInGoodStanding     := !user.IsBanned

isEligible := meetsAgeRequirement && isIdentityVerified && isInGoodStanding
if isEligible {
    // allow purchase
}
```

Each named boolean reads like a sentence and makes the combined condition self-documenting.

---

## 5. Repository Query Design

### Mandatory Split: Query vs Store Repository

Every domain must have **two** repository files with clearly separated responsibilities:

| File | Responsibility | Operations |
|------|---------------|------------|
| `{domain}_query_repository.go` | Read-only data access | `Find`, `Get`, `List`, `Exists`, `Count`, `Pagination` |
| `{domain}_store_repository.go` | Write operations | `Insert`, `Update`, `Delete`, `Save` |

This maps to two separate interfaces in `interfaces/`:
- `{domain}_query_repository_interface.go`
- `{domain}_store_repository_interface.go`

```go
// app/user/repositories/user_query_repository.go — read only
type UserQueryRepository struct{ db *gorm.DB }

func (r *UserQueryRepository) Pagination(ctx context.Context, q *infradtos.PaginationQueryRequestDTO) ([]UserListDTO, int64, error)
func (r *UserQueryRepository) FindUserByEmail(email string) (*entities.User, error)
func (r *UserQueryRepository) IsUserExistByEmail(email string) (bool, error)

// app/user/repositories/user_store_repository.go — write only
type UserStoreRepository struct{ db *gorm.DB }

func (r *UserStoreRepository) Insert(user *entities.User) error
func (r *UserStoreRepository) Update(user *entities.User) error
func (r *UserStoreRepository) SoftDelete(id uuid.UUID) error
```

### Query Naming Convention

Use purpose-specific names — no dynamic mega-functions.

Pattern: `Get + [What Data] + For + [Use Case] + By + [Filter]`

```go
// Bad: one dynamic function forced to serve all needs
func (r *UserRepository) GetUsers(status *string, city *string, roleID *int, includeOrders bool) ([]User, error)

// Good: each use case gets its own focused query
func (r *UserQueryRepository) Pagination(ctx context.Context, q *infradtos.PaginationQueryRequestDTO) ([]UserListDTO, int64, error)
func (r *UserQueryRepository) GetActiveUsersForDashboardByCity(city string) ([]UserDashboardDTO, error)
func (r *UserQueryRepository) GetUsersForGenerateReportByDateRange(start, end time.Time) ([]UserReportDTO, error)
func (r *UserQueryRepository) GetPendingUsersForVerificationByOlderThanDays(days int) ([]entities.User, error)
```

Each query returns a **DTO specific to its use case** — only load the fields needed.

### Pagination in Query Repository

For list endpoints the query repo method is named **`Pagination`**.
The page/offset calculation lives **inside** the repo via a `Paginate` scope helper.
The service layer passes the DTO — it does not compute offsets itself.

```go
// app/user/repositories/user_query_repository.go
func (r *UserQueryRepository) Pagination(
    ctx context.Context,
    q *infradtos.PaginationQueryRequestDTO,
) ([]UserListDTO, int64, error) {
    var results []UserListDTO
    var total int64

    query := r.db.WithContext(ctx).Model(&entities.User{})
    query = r.applyFilters(query, q)
    query = r.applySort(query, q)

    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    err := query.Scopes(utils.Paginate(q)).Select("id, name, email, status").Scan(&results).Error
    return results, total, err
}

// private helpers — not exposed on interface
func (r *UserQueryRepository) applyFilters(db *gorm.DB, q *infradtos.PaginationQueryRequestDTO) *gorm.DB {
    if q.Search != "" {
        db = db.Where("name ILIKE ?", q.Search+"%")
    }
    return db
}

func (r *UserQueryRepository) applySort(db *gorm.DB, q *infradtos.PaginationQueryRequestDTO) *gorm.DB {
    allowed := map[string]bool{"name": true, "created_at": true, "updated_at": true}
    col := q.SortBy
    if !allowed[col] {
        col = "created_at"
    }
    order := "DESC"
    if q.Order == enums.SortOrderAsc {
        order = "ASC"
    }
    return db.Order(col + " " + order)
}
```

**`Paginate` scope helper** lives in `infrastructure/utils/`:

```go
// infrastructure/utils/pagination_util.go
func Paginate(q *infradtos.PaginationQueryRequestDTO) func(*gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        page := q.Page
        if page <= 0 {
            page = 1
        }
        perPage := q.PerPage
        if perPage <= 0 || perPage > 100 {
            perPage = 10
        }
        return db.Offset((page - 1) * perPage).Limit(perPage)
    }
}
```

### Shared Pagination DTO

Pagination inputs live in **`infrastructure/dtos/`**, not inside any domain's `dtos/` folder.

```go
// infrastructure/enums/sort_order_enum.go
package enums

type SortOrderEnum string

const (
    SortOrderAsc  SortOrderEnum = "asc"
    SortOrderDesc SortOrderEnum = "desc"
)
```

```go
// infrastructure/dtos/pagination_query_request_dto.go
package dtos

import (
    "strconv"
    "github.com/gin-gonic/gin"
    "event-backend/infrastructure/enums"
)

type PaginationQueryRequestDTO struct {
    Search  string              `form:"search"`
    PerPage int                 `form:"per_page"`
    Page    int                 `form:"page"`
    SortBy  string              `form:"sort_by"`
    Order   enums.SortOrderEnum `form:"order"`
}

func NewPaginationQueryRequestDTO(c *gin.Context) *PaginationQueryRequestDTO {
    q := &PaginationQueryRequestDTO{}

    if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil && page >= 1 {
        q.Page = page
    } else {
        q.Page = 1
    }

    if perPage, err := strconv.Atoi(c.DefaultQuery("per_page", "10")); err == nil && perPage >= 1 && perPage <= 100 {
        q.PerPage = perPage
    } else {
        q.PerPage = 10
    }

    q.Search = c.Query("search")
    q.SortBy = c.DefaultQuery("sort_by", "created_at")

    order := enums.SortOrderEnum(c.DefaultQuery("order", string(enums.SortOrderDesc)))
    if order != enums.SortOrderAsc && order != enums.SortOrderDesc {
        order = enums.SortOrderDesc
    }
    q.Order = order

    return q
}
```

### Extending Per Domain — Struct Embedding

Each domain embeds `PaginationQueryRequestDTO` and adds its own filter fields.
`ShouldBindQuery` automatically binds all fields including embedded ones.

```go
// app/event/dtos/event_query_request_dto.go
package dtos

import (
    "github.com/google/uuid"
    infradtos "event-backend/infrastructure/dtos"
)

type EventQueryRequestDTO struct {
    infradtos.PaginationQueryRequestDTO       // embed base pagination

    CategoryID *uuid.UUID `form:"category_id"` // domain-specific filters
    Status     *string    `form:"status"`
    City       *string    `form:"city"`
}
```

**Usage in controller — bind with `ShouldBindQuery`, no manual constructor needed:**

```go
func (ctrl *EventController) Index(c *gin.Context) {
    var filter dtos.EventQueryRequestDTO
    if err := c.ShouldBindQuery(&filter); err != nil {
        // return 400
    }
    result, total, err := ctrl.eventService.GetEvents(c.Request.Context(), &filter)
    // ...
}
```

`filter.Page`, `filter.PerPage`, `filter.Search` are all directly accessible alongside `filter.CategoryID`, `filter.Status`.

**When there are NO extra filters**, use the base DTO directly:

```go
func (ctrl *UserController) Index(c *gin.Context) {
    var q infradtos.PaginationQueryRequestDTO
    if err := c.ShouldBindQuery(&q); err != nil { ... }
    result, total, err := ctrl.userService.GetUsers(c.Request.Context(), &q)
}
```

> `NewPaginationQueryRequestDTO` is kept for middleware or helpers that need manual default/validation logic outside `ShouldBindQuery`. For controllers, always prefer `ShouldBindQuery`.

**Usage in service:**

```go
func (s *EventService) GetEvents(ctx context.Context, q *dtos.EventQueryRequestDTO) ([]EventListDTO, int64, error) {
    return s.eventQueryRepo.Pagination(ctx, q)
}
```

---

## 6. Receiving Request Data in Gin

### Query Parameters (`?key=value`)

Use `c.Query()` / `c.DefaultQuery()` for manual parsing, or bind the whole struct with `c.ShouldBindQuery()`.

```go
// Manual
search := c.Query("search")           // "" if missing
page   := c.DefaultQuery("page", "1") // "1" if missing

// Struct bind (preferred for multi-param)
type FilterDTO struct {
    Search string `form:"search"`
    Status string `form:"status"`
    Page   int    `form:"page"`
}

var filter FilterDTO
if err := c.ShouldBindQuery(&filter); err != nil {
    // handle validation error
}
```

### Request Body (JSON)

Use `c.ShouldBindJSON()` with a DTO struct tagged `json:"..."`.

```go
type CreateEventDTO struct {
    Title      string    `json:"title"       validate:"required,min=3"`
    CategoryID uuid.UUID `json:"category_id" validate:"required"`
    StartAt    time.Time `json:"start_at"    validate:"required"`
}

var body CreateEventDTO
if err := c.ShouldBindJSON(&body); err != nil {
    // return 400
}
```

### Request Headers

Use `c.GetHeader()` for single headers.

```go
authHeader  := c.GetHeader("Authorization")  // "Bearer <token>"
contentType := c.GetHeader("Content-Type")
requestID   := c.GetHeader("X-Request-ID")
```

### Summary Table

| Source | Tag | Method |
|--------|-----|--------|
| Query param (`?page=1`) | `form:"page"` | `c.ShouldBindQuery(&dto)` or `c.Query("page")` |
| JSON body | `json:"title"` | `c.ShouldBindJSON(&dto)` |
| Header | — | `c.GetHeader("Authorization")` |
| Path param (`/users/:id`) | — | `c.Param("id")` |
| Form / multipart | `form:"file"` | `c.ShouldBind(&dto)` |

> Use `ShouldBind*` (not `Bind*`) — it does not abort the request automatically, giving you control over the error response.

---

## 7. Pointer Usage Best Practices

Use pointers deliberately — not by default.

| Situation | Use Pointer? | Reason |
|-----------|-------------|--------|
| Large struct passed to function | Yes | Avoid copying large data |
| Function needs to mutate the caller's value | Yes | Pass by reference |
| Optional field (can be nil/absent) | Yes | `nil` represents "not set" |
| Method needs to modify receiver state | Yes | Pointer receiver |
| Small, primitive-like struct (no mutation) | No | Value copy is cheap and safe |
| Return from repository (single entity) | Yes | `nil` signals "not found" cleanly |
| Return from repository (slice of entities) | No | `nil` slice and empty slice behave the same for ranging |

```go
// Pointer receiver — method mutates struct
func (s *UserService) Deactivate(id uuid.UUID) error { ... }

// Pointer return — nil means "not found"
func (r *UserQueryRepository) FindUserByEmail(email string) (*entities.User, error) {
    var user entities.User
    err := r.db.Where("email = ?", email).First(&user).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }
    return &user, err
}

// Pointer field — optional filter value
type EventQueryRequestDTO struct {
    infradtos.PaginationQueryRequestDTO
    Status    *string    `form:"status"`    // nil = no filter applied
    City      *string    `form:"city"`
    StartDate *time.Time `form:"start_date"`
}

// Bad: unnecessary pointer — small read-only value
func getPageSize(p *int) int { ... } // just pass int directly
```

---

## 8. Existence Checks

When you only need to know **whether a record exists**, never fetch the full model.
Use a dedicated `IsExist` method.

**Naming:** `Is + [Subject] + ExistBy + [Filter]`

```go
// Bad: fetches full row just to check existence
func (s *UserService) Register(email string) error {
    user, _ := s.userQueryRepo.FindUserByEmail(email)
    if user != nil {
        return errors.New("email already taken")
    }
    // ...
}

// Good: lightweight existence check in query repository
// app/user/repositories/user_query_repository.go
func (r *UserQueryRepository) IsUserExistByEmail(email string) (bool, error) {
    var exists bool
    err := r.db.Model(&entities.User{}).
        Select("1").
        Where("email = ? AND deleted_at IS NULL", email).
        Limit(1).
        Scan(&exists).Error
    return exists, err
}

// Caller in service
func (s *UserService) Register(req RegisterDTO) error {
    exists, err := s.userQueryRepo.IsUserExistByEmail(req.Email)
    if err != nil {
        return err
    }
    if exists {
        return ErrEmailAlreadyTaken
    }
    // ...
}
```

More examples:

```go
func (r *EventQueryRepository) IsEventExistByID(id uuid.UUID) (bool, error)
func (r *EventTicketQueryRepository) IsTicketExistByEventID(eventID uuid.UUID) (bool, error)
func (r *UserQueryRepository) IsUserExistByID(id uuid.UUID) (bool, error)
```

---

## 9. Avoid N+1 Queries

Use GORM `Preload` to eagerly load associations instead of querying in a loop.

```go
// Bad: N+1 — one extra query per event
events, _ := db.Find(&events)
for _, e := range events {
    db.Where("event_id = ?", e.ID).Find(&tickets)
}

// Good: 2 queries total
db.Preload("Tickets").Find(&events)

// Good for filtering on a relation: use Joins
db.Joins("JOIN event_tickets ON event_tickets.event_id = events.id AND event_tickets.quota > 0").
    Find(&events)
```

---

## 10. Avoid Monster Functions

A function should do one thing. If a function needs a comment to explain each block, split it.

```go
// Bad: one function doing everything
func (s *EventService) CreateEvent(req CreateEventDTO) error {
    // validate ...
    // check quota ...
    // create event ...
    // create tickets ...
    // send notification ...
    // invalidate cache ...
}

// Good: orchestrator delegates to focused helpers
func (s *EventService) CreateEvent(req CreateEventDTO) error {
    if err := s.validateCreateRequest(req); err != nil {
        return err
    }
    event := s.buildEventFromRequest(req)
    if err := s.eventStoreRepo.Insert(event); err != nil {
        return err
    }
    s.notifier.NotifyEventCreated(event)
    return nil
}
```
