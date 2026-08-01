# Backend Akademik

Sistem backend akademik yang dibangun dengan Go menggunakan clean architecture.

## Perintah Pengembangan

### Menjalankan Aplikasi
```bash
# Build aplikasi
go build -o tmp/main ./cmd/akademic

# Jalankan aplikasi
./tmp/main

# Atau gunakan Air untuk hot reload saat development
air
```

### Testing
```bash
# Jalankan semua test
go test ./...

# Jalankan test pada package tertentu
go test ./internal/repository/model/...

# Jalankan test dengan output verbose
go test -v ./tests/unit/...
```

### Docker
```bash
# Build Docker image
docker build -t akademik-be .

# Jalankan container Docker
docker run -p 9000:9000 akademik-be
```

## Struktur Project

```
cmd/akademic/       - Entry point aplikasi
internal/
  component/         - Dependency injection dan bootstrap
  config/            - Loading konfigurasi (Viper), database, cache, setup logger
  delivery/http/
    controller/      - Handler request HTTP
    middleware/      - Autentikasi, permission, error handling, metrics
    router/          - Definisi route per role user (student, lecturer, academic, programhead)
  dto/              - Data Transfer Objects dan converters
  model/            - Entitas database (model GORM)
  repository/
    model/          - Layer data access menggunakan stored procedure
    cached/         - Interface repository cache Redis
  service/
    model/          - Layer business logic
    command/         - Command handlers
pkg/
  auth/             - Verifikasi token JWT menggunakan RSA public key
  utils/            - Fungsi utility (generate UUID, error handling, dll)
  validation/       - Helper validasi request
tests/unit/         - Unit tests
```

## Arsitektur

### Dependency Injection
Semua dependency di-wire di fungsi Bootstrap `internal/component/app.go`. Repository, service, controller, dan middleware dibuat di sini dan diteruskan ke konfigurasi router.

### Stored Procedure Data Access
Codebase ini menggunakan raw SQL stored procedure via `db.Exec()` dan `db.Raw().Rows()`. Pola repository mengabstraksi ini:
- `internal/repository/model/*` - Berisi implementasi repository
- Repository menggunakan GORM ORM PostgreSQL queries (direfaktor dari legacy stored procedures)
- Konteks user diekstrak dari request context untuk field audit (created_by, updated_by, deleted_by)

### Cache Layer
Redis caching diimplementasi melalui interface `CacheRepository`. Repository menerima instance cache dan menggunakannya untuk data yang sering diakses.

### Integrasi Service Eksternal
Aplikasi terintegrasi dengan service internal UNSIA:
- **SDM** (Sumber Daya Manusia) - Endpoint general information
- **SSO** - Autentikasi dan permission checking
- **Datareferensi** - Data periode akademik

## Autentikasi & Autorisasi

### JWT Autentikasi
Verifikasi JWT berbasis RSA menggunakan public key (`public_key.pem`). Middleware autentikasi:
- Mengekstrak Bearer token dari header Authorization
- Verifikasi menggunakan `pkg/auth/token.go` → `VerifyTokenSpesifik()`
- Menyimpan `UserClaimsSpesifikRole` di request context sebagai `x-user-claims`

### Permission Middleware
`internal/delivery/http/middleware/permission_check.go` memvalidasi permission user via service permission SSO sebelum mengizinkan akses ke endpoint yang dilindungi.

## Konfigurasi

Konfigurasi diload via Viper dari environment variables atau file `.env`. Bagian konfigurasi utama:

| Bagian | Deskripsi |
|--------|------------|
| Database | Koneksi SQL Server dengan pooling |
| Server | Host, port, setting prefork |
| JWT | Path public key |
| Minio | Object storage untuk upload file |
| Redis | Layer caching |
| InternalService | Endpoint untuk SSO, SDM, Datareferensi |

## Pola Soft Delete

Aplikasi mengimplementasi soft delete menggunakan kolom timestamp `deleted_at` dan `deleted_by`. Record yang dihapus di-query terpisah via stored procedure `*HasDeleted` dan bisa di-restore.

## Prometheus Metrics

Aplikasi mengekspos metrik Prometheus di endpoint `/metrics` (dilindungi oleh `APP_KEY_METRICS`). Middleware custom melacak metrik request HTTP termasuk endpoint, method, dan status code.
