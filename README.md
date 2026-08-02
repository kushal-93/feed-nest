# FeedNest

A simple, lightweight RSS feed aggregator and user subscription backend service built in Go. This is a toy project designed for learning Go backend development, database integration, concurrency, API design, and clean handler decoration patterns for authentication.

## Features

- **User Management & Auth**: Create users and authenticate API requests using custom API keys.
- **Feed Subscription**: Users can add RSS feeds and follow/unfollow them.
- **Concurrent Background Scraper**: A background goroutine periodically fetches new posts from followed RSS feeds.
- **Post Feeds**: Fetch the latest posts aggregated from all the RSS feeds a user follows.

## Tech Stack

- **Language**: Go (v1.25.0)
- **Database**: PostgreSQL (v16)
- **Router**: Go-Chi (with CORS middleware)
- **SQL Code Generation**: sqlc (type-safe SQL compiler)
- **Migrations**: Goose

## Getting Started

### 1. Database Setup
Start the PostgreSQL database container:
```bash
docker-compose up -d
```

### 2. Environment Configuration
Create a `.env` file in the root directory (a sample is provided below):
```env
PORT=8080
DB_URL=postgres://postgres:password@localhost:5432/feednest?sslmode=disable
```

### 3. Database Migrations
Apply schema changes using `goose` (run from the root directory):
```bash
goose -dir sql/schema postgres "postgres://postgres:password@localhost:5432/feednest?sslmode=disable" up
```

### 4. Code Generation (Optional)
If you modify any SQL queries or schemas, regenerate the Go database package using `sqlc`:
```bash
sqlc generate
```

### 5. Running the Application
Build and run the Go backend:
```bash
go run .
```

## API Reference

All routes are prefixed with `/v1`.

| Method | Endpoint | Auth Required | Description |
| :--- | :--- | :--- | :--- |
| **GET** | `/healthz` | No | Health check endpoint |
| **GET** | `/error` | No | Test error handling response |
| **POST** | `/users` | No | Register a new user |
| **GET** | `/users` | Yes | Get the authenticated user details |
| **POST** | `/feeds` | Yes | Register a new RSS feed |
| **GET** | `/feeds` | No | List all registered RSS feeds |
| **POST** | `/feed_follows` | Yes | Subscribe to an RSS feed |
| **GET** | `/feed_follows` | Yes | Get all feed subscriptions for the user |
| **DELETE** | `/feed_follows/{id}` | Yes | Unsubscribe from an RSS feed |
| **GET** | `/posts` | Yes | Get posts from the user's subscribed feeds |

### Authentication

For endpoints requiring authentication, provide the user's API Key in the request header:
```http
Authorization: ApiKey <your_api_key>
```

#### Custom Authenticated Handler Pattern
The project uses a custom wrapper/decorator pattern (`middlewareAuth`) to handle authenticated routes:

```go
type authedHandler func(http.ResponseWriter, *http.Request, database.User)
```

This pattern extracts the API key, fetches the user, and injects the user details directly into downstream handlers. This avoids repetitive user authentication and database lookup code.
