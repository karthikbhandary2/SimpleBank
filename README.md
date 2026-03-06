# SimpleBank

A backend banking service built with Go, PostgreSQL, and gRPC that provides user management, account operations, and money transfers with email verification.

## Features

- User registration and authentication with JWT/PASETO tokens
- Email verification system with async task processing
- Account management (create, read, update)
- Money transfers between accounts with transaction support
- Role-based access control (depositor/banker)
- RESTful HTTP API and gRPC API
- Swagger documentation
- Redis-backed async task queue

## Tech Stack

- **Backend**: Go 1.x
- **Database**: PostgreSQL 18
- **Cache/Queue**: Redis 7
- **API**: gRPC with gRPC-Gateway for REST
- **Authentication**: JWT and PASETO tokens
- **Task Queue**: Asynq (Redis-based)
- **Database Migrations**: golang-migrate
- **Code Generation**: sqlc (SQL to Go), protoc (Protocol Buffers)
- **Testing**: Go testing with gomock
- **Frontend**: Vue 3 + TypeScript + Vite

## Prerequisites

- Go 1.19+
- Docker and Docker Compose
- Make
- migrate CLI tool
- sqlc
- protoc (Protocol Buffer compiler)

## Setup

### 1. Clone the repository

```bash
git clone <repository-url>
cd SimpleBank
```

### 2. Start services with Docker Compose

```bash
docker-compose up -d
```

This starts:
- PostgreSQL on port 5432
- Redis on port 6379
- API server on ports 8080 (HTTP) and 9090 (gRPC)

### 3. Manual setup (alternative)

If you prefer running services individually:

```bash
# Create Docker network
make network

# Start PostgreSQL
make postgres

# Create database
make createdb

# Run migrations
make migrateup

# Start Redis
make redis

# Run the server
make server
```

## Configuration

Edit `app.env` to configure:

```env
DB_SOURCE=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
REDIS_ADDRESS=localhost:6379
HTTP_SERVER_ADDRESS=0.0.0.0:8080
GRPC_SERVER_ADDRESS=0.0.0.0:9090
TOKEN_SYMMETRIC_KEY=<32-character-key>
ACCESS_TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=24h
ENVIRONMENT=development
EMAIL_SENDER_NAME=SimpleBank
EMAIL_SENDER_ADDRESS=<your-email>
EMAIL_SENDER_PASSWORD=<your-password>
```

## API Documentation

Once the server is running, access Swagger UI at:

```
http://localhost:8080/swagger/
```

## Database Schema

The database includes:

- **users**: User accounts with authentication
- **accounts**: Bank accounts with balances
- **entries**: Account balance change records
- **transfers**: Money transfer records between accounts
- **sessions**: User session management
- **verify_emails**: Email verification tokens

## Development

### Run tests

```bash
make test
```

### Generate code

```bash
# Generate Go code from SQL queries
make sqlc

# Generate gRPC/Protocol Buffer code
make proto

# Generate database mocks for testing
make mock
```

### Database migrations

```bash
# Create new migration
make new_migration name=<migration_name>

# Run all migrations
make migrateup

# Run one migration
make migrateup1

# Rollback all migrations
make migratedown

# Rollback one migration
make migratedown1
```

### Database documentation

```bash
# Generate database docs
make db_docs

# Generate SQL schema from DBML
make db_schema
```

## Project Structure

```
.
├── api/           # HTTP REST API handlers
├── db/
│   ├── migration/ # Database migration files
│   ├── query/     # SQL queries for sqlc
│   ├── sqlc/      # Generated Go code from SQL
│   └── mock/      # Mock database for testing
├── doc/           # Documentation and Swagger files
├── gapi/          # gRPC API handlers
├── mail/          # Email sending functionality
├── pb/            # Generated Protocol Buffer code
├── proto/         # Protocol Buffer definitions
├── token/         # JWT and PASETO token makers
├── util/          # Utility functions and config
├── val/           # Custom validators
├── worker/        # Async task processors
├── frontend/      # Vue.js frontend
└── main.go        # Application entry point
```

## API Endpoints

### HTTP (REST)
- `POST /users` - Create user
- `POST /users/login` - Login user
- `POST /accounts` - Create account
- `GET /accounts/:id` - Get account
- `POST /transfers` - Create transfer

### gRPC
- `CreateUser` - Register new user
- `LoginUser` - Authenticate user
- `UpdateUser` - Update user details
- `VerifyEmail` - Verify email address

## Testing

The project includes unit tests and integration tests using:
- Go's built-in testing package
- gomock for mocking dependencies
- testify for assertions

Run tests with coverage:

```bash
go test -v -cover ./...
```

## Deployment

The project includes:
- `Dockerfile` for containerization
- `docker-compose.yaml` for local deployment
- Kubernetes manifests in `eks/` directory
- GitHub Actions workflows in `.github/workflows/`
