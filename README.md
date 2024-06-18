# Hash File Service

## Overview

The Hash File Service is a backend service built using Nakama and Golang. The service provides an RPC function that processes file requests by reading files from disk, validating their JSON content, calculating their hash, and storing the file information in a PostgreSQL database. The service ensures that the file content is included in the response only if the provided hash matches the calculated hash.

## Installation and Setup

### Prerequisites

- Git
- Docker Compose

### Clone the Repository

```sh
git clone https://github.com/ShadeDS/hash-file-service.git
cd hash-file-service

docker compose up --build
```

## Testing the RPC Function
You can test the RPC function using the Nakama console or a curl request.

Example using curl:
```sh
curl -X POST "http://localhost:7350/v2/rpc/file_processing_rpc" \
     -d '{"type":"core","version":"1.0.0"}' \
     -H "Content-Type: application/json" \
     -H "Accept: application/json" \
     -H "Authorization: Bearer <NAKAMA_TOKEN>"
```

## Project Structure
```angular2html
hash-file-service/
├── database/
│   └── database.go
├── migrations/
│   ├── 000001_create_files_table.up.sql
│   └── 000001_create_files_table.down.sql
├── service/
│   └── service.go
├── util/
│   └── hash.go
├── main.go
├── Dockerfile
├── docker-compose.yml
├── local.yml
└── go.mod
```

## Explanation of the Solution
### Design Decisions
1. **Nakama**: Required by the description.
2. **Go**: Required by the description; the latest possible version was chosen for up-to-date features and updates.
3. **PostgreSQL**: Chosen because it's widely used and compatible with Nakama.
4. **Sha-256**: Chosen because it produces a 256-bit hash value, making it resistant to collisions.
5. **Golang-Migrate**: Chosen as it is a standard for database migrations.
6. **Testify** and **Go-sqlmock**: Chosen because these are the most popular libraries for unit testing.

## Thoughts and Ideas for Improvement
1. Nakama has its own mechanism for database migrations, but I didn't find the proper way to integrate my migrations with it.
2. For migrations, there is only up logic; there is no possibility to remove tables.
3. The configuration for the module should be externalized to support different environments.
4. The module contains only unit tests; it would be better to add integration tests as well.
5. All in all, it was a good opportunity to recall Golang.

