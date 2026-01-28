# beeBot Server

Enterprise-grade backend system based on Go and Fiber framework

## Project Structure

```
beeBot/
├── cmd/
│   └── server/
│       └── main.go          # Main entry point
├── config/
│   └── config.go           # Configuration management
├── handlers/               # HTTP request handlers
│   └── base_handlers.go
├── routes/                 # Route definitions
│   └── routes.go
├── models/                 # Data models
├── services/               # Business logic layer
├── utils/                  # Utility functions
├── api/                    # API definitions
├── go.mod
├── go.sum
└── Makefile
```

## Features

- High-performance web server based on Fiber framework
- Configuration management
- Health check endpoint
- Modular routing structure
- Scalable architecture

## Quick Start

### Requirements

- Go 1.18+

### Install Dependencies

```bash
go mod tidy
```

### Build Project

```bash
make build
```

### Run Server

```bash
make run
```

### Test Endpoints

- `GET /` - Welcome page
- `GET /health` - Health check
- `GET /api/v1/` - API endpoints

## Environment Variables

- `PORT` - Server port (default: 3000)
- `ENV` - Environment (default: development)
- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database user
- `DB_PASS` - Database password
- `DB_NAME` - Database name
- `JWT_SECRET` - JWT secret key

## Development Standards

- All code follows Go language standards
- Standard Go project structure
- Unified error handling
- Complete unit test coverage

## Deployment

```bash
# Production build
make build-linux

# Run
./bin/server-linux
```

## License

See LICENSE file
