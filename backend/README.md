# Entry Point Backend

A backend service for managing collections of cards and spaces, built with Go and using dependency injection with Wire.

## Getting Started

### Prerequisites
- Go 1.19 or higher
- PostgreSQL 14 or higher
- Wire (for dependency injection)

### Local Development Setup

1. Install dependencies and Wire:
```bash
go mod download
go install github.com/google/wire/cmd/wire@latest
```

2. Generate dependency injection code:
```bash
wire ./cmd
```

3. Configure Environment:
Copy a `.env.local` file to `.env` in the root directory and update the following variables:
```env
DB_URL="put_posgresql_url_here"
API_KEY=some_uuid
```

3. Run database migrations:
```bash
go run migrate/migrate.go
```

4. Generate Swagger documentation:
```bash
swag init -g cmd/main.go --parseDependency --parseInternal
```

5. Build the server:
```bash
go build -o main cmd/main.go cmd/wire_gen.go
```

5. Ruild the server:
```bash
./main
```

The API will be available at `http://localhost:8080`
Swagger documentation will be available at `http://localhost:8080/swagger/index.html`

## Run with docker compose
1. Prepare .env file

1. Run 
```bash
docker compose up
```

## Project Structure

### Dependency Injection
The project uses Google's [Wire](https://github.com/google/wire) for dependency injection. The wire configuration is located in:
- `cmd/wire.go` - Wire provider definitions
- `cmd/wire_gen.go` - Generated dependency injection code

Wire automatically manages the instantiation and injection of:
- Controllers
- Services
- Repositories
- Router setup

## Core System Objects

### Space
A Space is the top-level organizational unit that contains collections. It helps users organize their content into different areas or projects.

Key properties:
- ID
- Name
- User ID
- Created/Updated timestamps

### Collection
A Collection is a container for cards within a space. Collections help organize related cards together.

Key properties:
- ID
- Name
- Space ID
- Order (for sorting)
- Read-only flag
- Created/Updated timestamps

### Card
Cards are the basic content units that can be of different types:

#### Note Card
- Title
- Description
- Content (text)
- Collection ID
- Tags
- Search content

#### Link Card
- Title
- Description
- URL
- Collection ID
- Tags
- Search content

#### YouTube Video Card
- Title
- Description
- Video URL
- Transcript
- Collection ID
- Tags
- Search content

### Tag
Tags help categorize and search cards across collections and spaces.

Key properties:
- ID
- Name
- User ID

## API Documentation

The API is documented using Swagger. After starting the server, visit `http://localhost:3000/swagger/index.html` to view the interactive API documentation.

Main API endpoints:
- `/spaces` - Manage spaces
- `/collections` - Manage collections
- `/cards` - Manage cards
- `/cards/note` - Create note cards
- `/cards/link` - Create link cards
- `/cards/youtube` - Create YouTube video cards
- `/cards/search` - Search cards with filters