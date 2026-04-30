# Gin Ticketing API

A Go-based ticketing API built with Gin, MySQL, and SQLC.

## Development Setup

### Prerequisites
- Go 1.21+
- MySQL
- [Goose](https://github.com/pressly/goose) (for migrations)
- [SQLC](https://sqlc.dev/) (for type-safe SQL)
- [Air](https://github.com/cosmtrek/air) (for live reloading)

### Environment Variables
Create a `.env` file in the root directory:
```env
PORT=8899
DB_USER=root
DB_PASS=
DB_HOST=localhost
DB_PORT=3306
DB_NAME=go_ticketing
```

---

## Tutorial: Database & Development

### 1. Database Migrations (Goose)

Migrations are separated by module (e.g., `auth`, `tickets`).

#### Create a new migration
To create a new migration for the **auth** module:
```powershell
goose -dir internal/auth/sql/migrations create <migration_name> sql
```

#### Run migrations
To apply pending migrations to your local database:
```powershell
# For Auth module
goose -dir internal/auth/sql/migrations mysql "root:@tcp(localhost:3306)/go_ticketing" up

# For Tickets module
goose -dir internal/tickets/sql/migrations mysql "root:@tcp(localhost:3306)/go_ticketing" up
```

#### Check migration status
```powershell
goose -dir internal/auth/sql/migrations mysql "root:@tcp(localhost:3306)/go_ticketing" status
```

### 2. Generate Database Code (SQLC)

After creating migrations or updating `.sql` query files, generate the Go boilerplate:
```powershell
sqlc generate
```

### 3. Running the Application (Air)

Use **Air** for live reloading during development. It will automatically recompile and restart the server whenever you save a file.

```powershell
air
```

The API will be available at `http://localhost:8899`.
