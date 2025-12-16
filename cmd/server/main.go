
package main

import (
    "database/sql"
    "log"

    "github.com/gofiber/fiber/v2"
    _ "github.com/lib/pq"

    "go-user-api/db/sqlc"
    "go-user-api/internal/handler"
    "go-user-api/internal/logger"
    "go-user-api/internal/middleware"
    "go-user-api/internal/routes"
)

func main() {
    db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/usersdb?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }

    queries := sqlc.New(db)
    logg := logger.Init()

    app := fiber.New()
    app.Use(middleware.RequestID())
    app.Use(middleware.Logger(logg))

    userHandler := handler.NewUserHandler(queries)
    routes.Register(app, userHandler)

    log.Fatal(app.Listen(":8080"))
}
