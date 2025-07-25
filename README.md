# Authentication Service (SSO)

## Description
Authentication and authorization service (Single Sign-On) developed in Go. The project provides a user authentication system for various applications, implementing a mechanism for generating and verifying JWT tokens.

## Functionality
- Registration of new users
- User authentication
- Checking for administrator rights
- Support for various applications with their own secret keys
- Generation of JWT tokens for authorization

## Technologies
- Programming language: Go 1.24
- Database: SQLite3
- Communication: gRPC
- Authentication: JWT (JSON Web Tokens)
- Migrations: golang-migrate
- Configuration: cleanenv

## Project Structure
```
/cmd
  /migrator    # Utility for running migrations
  /sso         # Main application
/config        # Configuration files
/internal
  /app         # Entry point to the application
  /config      # Configuration management
  /domain      # Domain models
  /grpc        # gRPC servers and handlers
  /lib         # Helper libraries
    /jwt       # Working with JWT tokens
    /logger    # Logging
  /services    # Business logic
    /auth      # Authentication service
  /storage     # Working with data storage
    /sqlite    # Implementation for SQLite
/migrations    # SQL migrations for the database
/storage       # Directory for storing DB files
```

## Running the Application

### Prerequisites
- Go 1.24 or newer
- CGO must be enabled for SQLite to work (CGO_ENABLED=1)

### Starting the Service
1. Make sure your configuration file is set up correctly:
```bash
export CONFIG_PATH=./config/local.yaml
# or pass the config path via flag
go run ./cmd/sso/main.go --config=./config/local.yaml
```

2. Running migrations:
```bash
go run ./cmd/migrator --storage-path=./storage/sso.db --migrations-path=./migrations
```

3. Starting the service:
```bash
go run ./cmd/sso/main.go
```

## Configuration
Example configuration file (`config/local.yaml`):
```yaml
env: "local"
storage_path: "./storage/sso.db"
token_ttl: "1h"
grpc:
  port: 44044
  timeout: "5s"
```

## Migrations
SQL migrations are used in the project to initialize and update the database structure. Migrations are located in the `/migrations` directory.

### List of migrations:
1. `1_init.up.sql` - Creating main tables (users, apps)
2. `1_init.down.sql` - Dropping tables
3. `2_add_is_admin_column_to_users_tbl.up.sql` - Adding is_admin column
4. `2_add_is_admin_column_to_users_tbl.down.sql` - Removing is_admin column

## Development Notes
- For development on Windows, make sure to set CGO_ENABLED=1 when compiling
- Use `go mod tidy` to manage dependencies
- For testing gRPC interfaces, it's recommended to use tools like grpcurl or BloomRPC
