# Dating Wars Backend

A modern dating application backend built with Go.

## Project Structure

```
.
├── cmd/                  # Main applications
│   └── api/              # API server
├── internal/             # Private application and library code
│   ├── config/           # Configuration management
│   ├── database/         # Database related code
│   ├── handler/          # HTTP handlers
│   ├── middleware/       # HTTP middleware
│   ├── model/            # Data models
│   ├── repository/       # Data access layer
│   └── service/          # Business logic
├── pkg/                  # Public library code
├── api/                  # API documentation and OpenAPI specs
├── migrations/           # Database migrations
└── scripts/              # Build and deployment scripts
```

## Features (Planned)

- User authentication and authorization
- Profile management
- Matching system
- Messaging system
- Location-based features
- Preferences and settings
- Admin dashboard

## Getting Started

1. Clone the repository
2. Install dependencies: `go mod download`
3. Set up environment variables
4. Run the server: `go run cmd/api/main.go`

## Development

- Go 1.24.0 or higher
- PostgreSQL (for database)
- Redis (for caching)

## License

See [LICENSE](LICENSE) file for details.