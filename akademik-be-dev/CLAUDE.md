# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Running the Application
```bash
# Build the application
go build -o tmp/main ./cmd/akademic

# Run the application
./tmp/main

# Or use Air for hot reload during development
air
```

### Testing
```bash
# Run all tests
go test ./...

# Run tests in a specific package
go test ./internal/repository/model/...

# Run tests with verbose output
go test -v ./tests/unit/...
```

### Docker
```bash
# Build Docker image
docker build -t akademik-be .

# Run Docker container
docker run -p 9000:9000 akademik-be
```

### Database Migrations
The application uses stored procedures in SQL Server for database operations. Stored procedures follow naming pattern:
- `sp_mst_{entity}_insert` - Create operations
- `sp_mst_{entity}_update_by_id` - Update operations
- `sp_mst_{entity}_delete_by_id` - Soft delete operations
- `sp_mst_{entity}_restore_by_id` - Restore soft deleted records
- `sp_mst_{entity}_get` - List with pagination
- `sp_mst_{entity}_get_by_id` - Get single record

## Architecture Overview

This is an academic backend system built with Go following a clean architecture pattern.

### Project Structure
```
cmd/akademic/       - Application entry point
internal/
  component/         - Dependency injection and bootstrap
  config/            - Configuration loading (Viper), database, cache, logger setup
  delivery/http/
    controller/      - HTTP request handlers
    middleware/      - Authentication, permissions, error handling, metrics
    router/          - Route definitions grouped by user role (student, lecturer, academic, programhead)
  dto/              - Data Transfer Objects and converters
  model/            - Database entities (GORM models)
  repository/
    model/          - Data access layer using stored procedures
    cached/         - Redis cache repository interface
  service/
    model/          - Business logic layer
    command/         - Command handlers
pkg/
  auth/             - JWT token verification using RSA public key
  utils/            - Utility functions (UUID generation, error handling, etc.)
  validation/       - Request validation helpers
tests/unit/         - Unit tests
```

### Key Architecture Patterns

**Dependency Injection**: All dependencies are wired in `internal/component/app.go` Bootstrap function. Repositories, services, controllers, and middleware are created here and passed into the router configuration.

**Stored Procedure Data Access**: Unlike typical GORM usage, this codebase uses raw SQL stored procedures via `db.Exec()` and `db.Raw().Rows()`. The repository pattern abstracts this:
- `internal/repository/model/*` - Contains repository implementations
- Repositories use GORM ORM PostgreSQL queries (refactored from legacy stored procedures)
- User context is extracted from request context for audit fields (created_by, updated_by, deleted_by)

**Cache Layer**: Redis caching is implemented through the `CacheRepository` interface. Repositories accept a cache instance and use it for frequently accessed data.

**External Service Integration**: The app integrates with internal UNSIA services:
- SDM (Sumber Daya Manusia) - General information endpoints
- SSO - Authentication and permission checking
- Datareferensi - Academic period data

### Authentication & Authorization

**JWT Authentication**: RSA-based JWT verification using public key (`public_key.pem`). The authentication middleware:
- Extracts Bearer token from Authorization header
- Verifies using `pkg/auth/token.go` → `VerifyTokenSpesifik()`
- Stores `UserClaimsSpesifikRole` in request context as `x-user-claims`

**Permission Middleware**: `internal/delivery/http/middleware/permission_check.go` validates user permissions via the SSO permission service before allowing access to protected endpoints.

### Request/Response Handling

**Pagination**: Uses `internal/dto/pageable/PageableRequestClass` for list endpoints. Default pagination parameters are applied via `GetDefault*()` helper methods.

**Validation**: `pkg/validation/validation.go` provides a generic `Validate()` function that processes struct validation errors and returns formatted error messages.

**Error Handling**: All controllers follow the same error handling pattern:
1. Extract `code-error` from locals (tracking identifier)
2. Parse request body/query
3. Validate request
4. Call service layer
5. On error: log with context and return appropriate status
6. On success: return success response

**Response Format**: Uses `dto.CreateSuccess()` and `dto.CreateError()` for consistent JSON responses.

### Configuration

Configuration is loaded via Viper from environment variables or `.env` file. Key configuration sections:
- Database: SQL Server connection with pooling
- Server: Host, port, prefork settings
- JWT: Public key path
- Minio: Object storage for file uploads
- Redis: Caching layer
- InternalService: Endpoints for SSO, SDM, Datareferensi

### Soft Delete Pattern

The application implements soft delete using `deleted_at` and `deleted_by` timestamp columns. Deleted records are queried separately via `*HasDeleted` stored procedures and can be restored.

### User Context Management

User authentication context flows through the request:
1. Auth middleware stores `UserClaimsSpesifikRole` at `x-user-claims` key
2. Controllers retrieve via `ctx.Locals("x-user-claims")`
3. Services access via `middleware.GetUserClaimsCtx(ctx)` helper
4. Repositories extract from context for audit fields

### File Upload Handling

File uploads use multipart form data. Controllers call `ctx.FormFile()` to extract files, then pass to `StorageService` for Minio upload. The `dto.MstClassContractRequest` pattern shows how to handle file fields alongside other form data.

### Prometheus Metrics

The application exposes Prometheus metrics at `/metrics` endpoint (protected by `APP_KEY_METRICS`). Custom middleware tracks HTTP request metrics including endpoint, method, and status code.
