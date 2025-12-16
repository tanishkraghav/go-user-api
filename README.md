
# Go User API

## Setup
1. Create PostgreSQL DB `usersdb`
2. Run migrations from `db/migrations`
3. Run `sqlc generate` inside `db/sqlc`
4. Start server:
   go run cmd/server/main.go

## Endpoints
POST /users
GET /users
GET /users/:id
PUT /users/:id
DELETE /users/:id
