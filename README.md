# go-crud

A minimal CRUD REST API for managing clinics, built with Go, Gin and GORM (Postgres).

This repository implements a simple generic CRUD service and example Clinic model. It demonstrates:

- Generic service wrappers for list/get/create/update/delete operations
- Pagination and simple search (query param `q`)
- Gin for HTTP routing
- GORM with the Postgres driver for persistence
- Environment configuration with `godotenv`

## Tech stack

- Go (module: `github.com/sithuramyo/go-crud`)
- Gin (github.com/gin-gonic/gin)
- GORM + Postgres (gorm.io/gorm, gorm.io/driver/postgres)

## Requirements

- Go (tested with the module's `go` directive, see `go.mod`)
- PostgreSQL server

## Environment

Create a `.env` file in the project root with the following variables (example):

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=Rootroot1
DB_NAME=go_crud

Note: `initializers/load_env.go` uses `github.com/joho/godotenv` to load `.env`.

## Database

The project connects to Postgres in `initializers/database.go`. The current code establishes a connection but does not run automatic schema migrations. You have two options:

1) Create the `clinics` table manually in Postgres.
2) Add an AutoMigrate call to `initializers/database.go` (example):

```go
// after successful gorm.Open...
DB.AutoMigrate(&models.Clinic{})
```

The `Clinic` model is defined in `models/clinic.go` and uses a UUID primary key generated in `BeforeCreate`.

## Build & Run

Run the server locally (PowerShell example):

```powershell
# build
go build -o go-crud.exe .
# run
.\\go-crud.exe
# or run directly with go run
go run .
```

The server will start on the default Gin port (8080). You can change this in `main.go` by passing an address to `r.Run(":8000")`.

## API Endpoints

Base path: `/api/clinics`

- GET `/api/clinics` — list clinics (supports `page`, `limit`, `q` search)
- POST `/api/clinics` — create a clinic
- GET `/api/clinics/:id` — get single clinic
- PUT `/api/clinics/:id` — update clinic
- DELETE `/api/clinics/:id` — delete clinic

Sample curl requests:

```bash
# List (first page)
curl http://localhost:8080/api/clinics

# Search and paginate
curl "http://localhost:8080/api/clinics?q=health&page=2&limit=5"

# Create
curl -X POST http://localhost:8080/api/clinics -H "Content-Type: application/json" -d '{"name":"Good Clinic","address":"123 Main St","contactNumber":"123-456"}'

# Get
curl http://localhost:8080/api/clinics/<uuid>

# Update
curl -X PUT http://localhost:8080/api/clinics/<uuid> -H "Content-Type: application/json" -d '{"name":"Updated"}'

# Delete
curl -X DELETE http://localhost:8080/api/clinics/<uuid>
```

## Request & Response shapes

The `Clinic` model fields (JSON keys):

- `id` (uuid)
- `name` (string)
- `address` (string)
- `contactNumber` (string)
- `createdAt` (timestamp)
- `updatedAt` (timestamp)
- `deletedAt` (nullable timestamp)

List responses are returned as JSON with pagination metadata: `data`, `page`, `limit`, `total`, `totalPages`.

## Notes & Next steps

- The DB connection currently uses a hardcoded DSN in `initializers/database.go`; it's recommended to use the env variables there. Example DSN construction:

```go
dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	host, user, password, dbname, port)
```

- Consider adding `DB.AutoMigrate(&models.Clinic{})` if you want GORM to create the table automatically.
- Add validation for input payloads (e.g., `github.com/go-playground/validator/v10`).
- Add tests for service wrappers and controllers.

## License

This repository does not include a license file. Add a LICENSE if you plan to open-source it.

----

If you'd like, I can update `initializers/database.go` to build the DSN from env vars and optionally call `AutoMigrate` — should I make those changes now?
