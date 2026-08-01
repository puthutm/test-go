---
name: create-endpoint
description: Generate scaffold code for new API endpoints following Clean Architecture pattern. Use when creating new CRUD endpoints, toggle endpoints, or paginated list endpoints in the akademik-be backend.
argument-hint: entity-name
disable-model-invocation: true
---

# Create New Endpoint - Akademik Backend

This skill generates boilerplate code for new API endpoints following the Clean Architecture pattern used in this akademik-be project.

## Interactive Prompt

Follow these steps to create a new endpoint:

### Step 1: Basic Information

First, ask the user for:
- **Entity name** (PascalCase): e.g., `OpenCloseValue`, `ClassSchedule`
- **Table name**: Database table name (e.g., `trx_open_close_values`, `mst_class_schedules`)
- **Description**: Brief description of what this entity does

### Step 2: HTTP Methods to Generate

Ask which HTTP methods to generate (user can select multiple):
- **GET List** - Get list of records with pagination/filtering
- **GET Detail** - Get single record by ID
- **PUT** - Update/toggle existing record
- **POST** - Create new record
- **DELETE** - Soft delete record

### Step 3: Route Configuration

Ask for:
- **Route path**: e.g., `/api/academic/setting/academic-period/:academic_periode_id/classes/open-close-values`
- **Route group**: Which router file to modify (`academic.go`, `programhead.go`, `lecturer.go`, `student.go`)
- **Role access**: Which roles can access this endpoint (academic, programhead, lecturer, student)

### Step 4: Stored Procedure Names

For each selected method, ask for the full stored procedure definition or name. Common patterns:
- `sp_mst_{entity}_insert` - Create
- `sp_mst_{entity}_update_by_id` - Update
- `sp_mst_{entity}_delete_by_id` - Delete
- `sp_mst_{entity}_get` - List with pagination
- `sp_mst_{entity}_get_by_id` - Get single record
- `sp_trx_{entity}_{action}` - Transaction operations

If the user provides the full SP definition, extract parameter names and types from it.

### Step 5: Request/Response Fields

Ask for field definitions based on the stored procedure parameters:
- **Field name** (camelCase for Go)
- **Go type**: `string`, `int`, `int64`, `bool`, `uuid.UUID`, or pointer types for nullable fields
- **Validation rules**: `required`, `uuid`, `oneof=X Y`, etc.
- **JSON tag**: How it appears in request/response body
- **Source**: `body` (from request body), `query` (from query params), `param` (from URL params), or `form` (for multipart)

## File Generation

After collecting information, generate the following files:

### 1. DTO File
Location: `internal/dto/{entity_snake}_dto.go`

**For simple toggle/update endpoints:**
```go
package dto

type {Entity}Request struct {
    FieldFromParam string `json:"-"`
    FieldFromBody  bool   `json:"field_from_body"`
}
```

**For paginated list endpoints:**
- Use existing `pageable.PageableRequestClass` or create custom pageable struct in `internal/dto/pageable/`
- Response should use `pageable.PageableResponse[dto.{Entity}Response]`

**Important:** Do NOT add `validate:"required"` on `bool` fields — `false` will fail required validation. Use `validate:"required"` only on non-boolean types.

**Note:** For entities that reuse existing models (e.g., `MstClass`), do NOT create a new model or DTO — reuse the existing ones with their converters.

### 2. Repository File (Interface Pattern)
Location: `internal/repository/model/{entity_snake}_repository.go`

```go
package repositorymodel

import (
    "context"
    "time"

    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
    "unsia.ac.id/akademic_be/internal/dto"
    "unsia.ac.id/akademic_be/pkg/auth"
)

type {Entity}Repository interface {
    MethodA(db *gorm.DB, ctx context.Context, req dto.{Entity}Request) error
    MethodB(db *gorm.DB, pageable pageable.PageableRequest{Custom}) (T []model.{Model}, count int64, err error)
}

type {entityCamel}Repository struct {
    log *logrus.Logger
}

func New{Entity}Repository(
    log *logrus.Logger,
) {Entity}Repository {
    return &{entityCamel}Repository{
        log: log,
    }
}

func (r *{entityCamel}Repository) MethodA(db *gorm.DB, ctx context.Context, req dto.{Entity}Request) error {
    user := ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)

    query := `EXEC dbo.sp_{sp_name}
        @param1 = ?,
        @param2 = ?,
        @created_at = ?,
        @created_by = ?
    `
    return db.Exec(query,
        req.Param1,
        req.Param2,
        time.Now().UnixMilli(),
        user.ID,
    ).Error
}
```

**For paginated list SPs (two result sets):**
```go
func (r *{entityCamel}Repository) GetAllWithCount(db *gorm.DB, pageable pageable.PageableRequestClass) (T []model.{Model}, count int64, err error) {
    query := `EXEC dbo.sp_{sp_name}
        @Filter = ?, @SortBy = ?, @SortDirection = ?,
        @Page = ?, @PageSize = ?,
        @AcademicPeriodeId = ?
    `

    rows, err := db.Raw(query,
        pageable.GetDefaultSearch(),
        pageable.GetDefaultSortBy(),
        pageable.GetDefaultSort(),
        pageable.GetDefaultPage(),
        pageable.GetDefaultLimit(),
        pageable.GetDefaultAcademicPeriodeId(),
    ).Rows()
    if err != nil {
        return T, count, err
    }
    defer rows.Close()

    for rows.Next() {
        var data model.{Model}
        if err := db.ScanRows(rows, &data); err != nil {
            return T, count, err
        }
        T = append(T, data)
    }

    if err := rows.Err(); err != nil {
        return T, count, err
    }

    // Second result set: total count
    if rows.NextResultSet() {
        if rows.Next() {
            if err := rows.Scan(&count); err != nil {
                return T, count, err
            }
        }
    }

    return T, count, err
}
```

### 3. Service File (Interface Pattern)
Location: `internal/service/model/{entity_snake}_service.go`

```go
package servicemodel

import (
    "context"

    "github.com/sirupsen/logrus"
    "gorm.io/gorm"
    "unsia.ac.id/akademic_be/internal/dto"
    "unsia.ac.id/akademic_be/internal/dto/converter"
    "unsia.ac.id/akademic_be/internal/dto/pageable"
    repositorymodel "unsia.ac.id/akademic_be/internal/repository/model"
    "unsia.ac.id/akademic_be/pkg/utils"
)

type {Entity}Service interface {
    MethodA(ctx context.Context, req dto.{Entity}Request) error
    GetAllWithCount(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.{Entity}Response], error)
}

type {entityCamel}Service struct {
    log                *logrus.Logger
    db                 *gorm.DB
    {entityCamel}Repo  repositorymodel.{Entity}Repository
}

func New{Entity}Service(
    log *logrus.Logger,
    db *gorm.DB,
    {entityCamel}Repo repositorymodel.{Entity}Repository,
) {Entity}Service {
    return &{entityCamel}Service{
        log:               log,
        db:                db,
        {entityCamel}Repo: {entityCamel}Repo,
    }
}

func (s *{entityCamel}Service) MethodA(ctx context.Context, req dto.{Entity}Request) error {
    tx := s.db.WithContext(ctx)

    err := s.{entityCamel}Repo.MethodA(tx, ctx, req)
    if err != nil {
        createMsg := utils.CreateMsgDebuging(err.Error(), req.ID, utils.ErrorLocation())
        s.log.WithError(err).Error(createMsg)
        return utils.ErrorSpToErrorFiber(err)
    }

    return nil
}
```

**For paginated list methods:**
```go
func (s *{entityCamel}Service) GetAllWithCount(ctx context.Context, pageble pageable.PageableRequestClass) (*pageable.PageableResponse[dto.{Entity}Response], error) {
    tx := s.db.WithContext(ctx)

    data, totalData, err := s.{entityCamel}Repo.GetAllWithCount(tx, pageble)
    if err != nil {
        createMsg := utils.CreateMsgDebuging(err.Error(), "", utils.ErrorLocation())
        s.log.WithError(err).Error(createMsg)
        return nil, utils.ErrorSpToErrorFiber(err)
    }

    res := make([]dto.{Entity}Response, 0, totalData)
    for _, item := range data {
        res = append(res, *converter.{Entity}ModelToResponse(item))
    }

    return &pageable.PageableResponse[dto.{Entity}Response]{
        Data: res,
        Metadata: pageable.Metadata{
            TotalData: totalData,
            TotalPage: utils.TotalPage(totalData, pageble.GetDefaultLimit()),
            Page:      pageble.GetDefaultPage(),
            Size:      pageble.GetDefaultLimit(),
        },
    }, nil
}
```

### 4. Controller File
Location: `internal/delivery/http/controller/{entity_snake}_controller.go`

```go
package controller

import (
    "fmt"

    "github.com/go-playground/validator/v10"
    "github.com/gofiber/fiber/v2"
    "github.com/sirupsen/logrus"
    msg "unsia.ac.id/akademic_be/internal/config/message"
    "unsia.ac.id/akademic_be/internal/dto"
    "unsia.ac.id/akademic_be/internal/dto/pageable"
    servicemodel "unsia.ac.id/akademic_be/internal/service/model"
    "unsia.ac.id/akademic_be/pkg/auth"
    "unsia.ac.id/akademic_be/pkg/utils"
    "unsia.ac.id/akademic_be/pkg/validation"
)

type {Entity}Controller struct {
    log            *logrus.Logger
    {entityCamel}Service servicemodel.{Entity}Service
    validate       *validator.Validate
}

func New{Entity}Controller(
    log *logrus.Logger,
    {entityCamel}Service servicemodel.{Entity}Service,
    validate *validator.Validate,
) *{Entity}Controller {
    return &{Entity}Controller{
        log:            log,
        {entityCamel}Service: {entityCamel}Service,
        validate:       validate,
    }
}
```

**For toggle/update handlers:**
```go
func (c *{Entity}Controller) UpdateByX(ctx *fiber.Ctx) error {
    codeError := ctx.Locals("code-error").(string)
    user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)
    paramID := ctx.Params("param_id")

    var req dto.{Entity}Request
    if err := ctx.BodyParser(&req); err != nil {
        utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
            "code-error":          codeError,
            "user":                user.ID,
            "location":            utils.ErrorLocation(),
            "message-error-debug": fmt.Sprintf("Failed to parse body : %+v", err),
        })

        statusCode := fiber.StatusInternalServerError
        if fiberErr, ok := err.(*fiber.Error); ok {
            statusCode = fiberErr.Code
        }

        return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
    }

    req.ParamID = paramID

    fails := validation.Validate(c.validate, &req)
    if len(fails) > 0 {
        return ctx.Status(fiber.StatusUnprocessableEntity).JSON(dto.CreateErrorValidation(fails))
    }

    err := c.{entityCamel}Service.UpdateByX(ctx.Context(), req)
    if err != nil {
        utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
            "code-error":          codeError,
            "user":                user.ID,
            "location":            utils.ErrorLocation(),
            "message-error-debug": fmt.Sprintf("Failed to update {entity} by ID : %s", paramID),
        })

        statusCode := fiber.StatusInternalServerError
        if fiberErr, ok := err.(*fiber.Error); ok {
            statusCode = fiberErr.Code
        }

        return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
    }

    return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessUpdate, ""))
}
```

**For paginated list handlers:**
```go
func (c *{Entity}Controller) GetAllWithCount(ctx *fiber.Ctx) error {
    codeError := ctx.Locals("code-error").(string)
    user := ctx.Locals("x-user-claims").(*auth.UserClaimsSpesifikRole)

    // If URL params are needed
    paramID := ctx.Params("academic_periode_id")
    paramID2 := ctx.Params("study_program_id")

    var pageble pageable.PageableRequestClass
    if err := ctx.QueryParser(&pageble); err != nil {
        utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
            "code-error":          codeError,
            "user":                user.ID,
            "location":            utils.ErrorLocation(),
            "message-error-debug": fmt.Sprintf("Failed to parse query : %+v", err),
        })

        statusCode := fiber.StatusInternalServerError
        if fiberErr, ok := err.(*fiber.Error); ok {
            statusCode = fiberErr.Code
        }

        return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
    }

    _, err := utils.ValidateAndPrepareRequest(&pageble)
    if err != nil {
        return err
    }

    // Set params from URL into pageable
    pageble.SetDefaultAcademicPeriodeId(paramID)
    pageble.SetDefaultStudyProgramId(paramID2)

    data, err := c.{entityCamel}Service.GetAllWithCount(ctx.Context(), pageble)
    if err != nil {
        utils.CreateCaptureAndLogFileError(c.log, err, map[string]any{
            "code-error":          codeError,
            "user":                user.ID,
            "location":            utils.ErrorLocation(),
            "message-error-debug": "Failed to get data",
        })

        statusCode := fiber.StatusInternalServerError
        if fiberErr, ok := err.(*fiber.Error); ok {
            statusCode = fiberErr.Code
        }

        return ctx.Status(statusCode).JSON(dto.CreateError(statusCode, codeError, err.Error()))
    }

    return ctx.Status(fiber.StatusOK).JSON(dto.CreateSuccess(fiber.StatusOK, msg.SuccessRead, data))
}
```

## Automatic Modifications

### Modify `internal/delivery/http/router/route.go`

Add controller field to `RouterConfig` struct:
```go
{Entity}Controller *controller.{Entity}Controller
```

### Modify `internal/component/app.go`

Add DI wiring in the correct sections (follow existing patterns):

**Repository section** (after existing repositories):
```go
{entityCamel}Repository := repositorymodel.New{Entity}Repository(
    component.Log,
)
```

**Service section** (after existing services):
```go
{entityCamel}Service := servicemodel.New{Entity}Service(
    component.Log,
    component.DB,
    {entityCamel}Repository,
)
```

**Controller section** (after existing controllers):
```go
{entityCamel}Controller := controller.New{Entity}Controller(
    component.Log,
    {entityCamel}Service,
    component.Validate,
)
```

**RouterConfig struct** (add controller field):
```go
{Entity}Controller: {entityCamel}Controller,
```

### Modify Router File

Add routes in the appropriate router file (`internal/delivery/http/router/{role}.go`):

```go
// In the appropriate route group section
{entityGroup} := parentGroup.Group("/path")
{entityGroup}.Get("/", r.{Entity}Controller.GetAllWithCount)
{entityGroup}.Put("/", r.{Entity}Controller.UpdateByX)
```

**Important:** When using URL params like `:academic_periode_id`, make sure the route group path includes the param:
```go
openCloseValues := academicPeriod.Group("/:academic_periode_id/classes/open-close-values")
```

Literal path segments must be defined before parametric routes to avoid conflicts.

## Reusing Existing Layers

When the new endpoint operates on an existing entity (e.g., adding a new GET method for `MstClass`):

- **Do NOT create new files** — add methods to the existing repository, service, and controller
- Add the new method to the existing interface definition
- Add the route to the appropriate router file
- No changes needed in `app.go` since the controller already exists

## Interface Convention (MANDATORY)

Repository and Service layers MUST use the interface pattern:

### Repository Interface
```go
type {Entity}Repository interface {
    Create(db *gorm.DB, ctx context.Context, req dto.{Entity}Request) error
    GetAllWithCount(db *gorm.DB, ...) ([]model.{Entity}, int64, error)
}

type {entityCamel}Repository struct {
    log *logrus.Logger
}

func New{Entity}Repository(log *logrus.Logger) {Entity}Repository {
    return &{entityCamel}Repository{log: log}
}
```

### Service Interface
```go
type {Entity}Service interface {
    Create(ctx context.Context, req dto.{Entity}Request) error
    GetAllWithCount(ctx context.Context, ...) (*pageable.PageableResponse[dto.{Entity}Response], error)
}

type {entityCamel}Service struct {
    log  *logrus.Logger
    db   *gorm.DB
    repo repositorymodel.{Entity}Repository
}

func New{Entity}Service(log *logrus.Logger, db *gorm.DB, repo repositorymodel.{Entity}Repository) {Entity}Service {
    return &{entityCamel}Service{log: log, db: db, repo: repo}
}
```

### Controller accepts Service interface
```go
type {Entity}Controller struct {
    log      *logrus.Logger
    service  servicemodel.{Entity}Service
    validate *validator.Validate
}

func New{Entity}Controller(
    log *logrus.Logger,
    service servicemodel.{Entity}Service,
    validate *validator.Validate,
) *{Entity}Controller { ... }
```

### Rules
- Interface: exported name (`MstClassRepository`, `MstClassService`)
- Struct: unexported name (`mstClassRepository`, `mstClassService`)
- Constructor: returns interface, not struct pointer
- Controller accepts service interface, service accepts repository interface

## Reference Files

Refer to these existing files for implementation patterns:

**Simple toggle endpoint (trx_open_close_value):**
- `internal/dto/trx_open_close_value_dto.go` - DTO
- `internal/repository/model/trx_open_close_value_repository.go` - Repository with interface
- `internal/service/model/trx_open_close_value_service.go` - Service with interface
- `internal/delivery/http/controller/trx_open_close_value_controller.go` - Controller

**Paginated list endpoint (mst_class):**
- `internal/dto/mst_class_dto.go` - DTO with response struct
- `internal/dto/converter/mst_class_converter.go` - Model to response converter
- `internal/dto/pageable/pageable_class.go` - Pageable request with custom fields
- `internal/repository/model/mst_class_repository.go` - Repository with two-result-set SP call
- `internal/service/model/mst_class_service.go` - Service with PageableResponse
- `internal/delivery/http/controller/mst_class_controller.go` - Controller with QueryParser

**DI wiring:**
- `internal/component/app.go` - Bootstrap function
- `internal/delivery/http/router/route.go` - RouterConfig struct
- `internal/delivery/http/router/academic.go` - Academic routes

## Important Notes

1. **Always use stored procedures** for database operations — never inline SQL queries in service/controller
2. **UUID fields** use `char(36)` GORM type mapping
3. **Nullable fields** use pointer types (`*string`, `*int64`)
4. **Timestamps** use `int64` (Unix milliseconds via `time.Now().UnixMilli()`), not `time.Time`
5. **User context** always extract via `ctx.Value("x-user-claims").(*auth.UserClaimsSpesifikRole)` in repository
6. **Error handling** use `utils.ErrorSpToErrorFiber()` for SQL/SP errors in service
7. **Error logging** use `utils.CreateCaptureAndLogFileError()` in controller, `utils.CreateMsgDebuging()` in service
8. **Validation** use `validation.Validate(c.validate, &req)` in controller
9. **Responses** use `dto.CreateSuccess()` and `dto.CreateError()` for consistent formatting
10. **Interface pattern** MUST be used for Repository and Service layers
11. **No `validate:"required"` on bool fields** — `false` fails required validation
12. **Run `go build`** after all changes to verify compilation
