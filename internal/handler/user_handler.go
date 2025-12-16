
package handler

import (
    "strconv"

    "github.com/gofiber/fiber/v2"
    "github.com/go-playground/validator/v10"

    "go-user-api/db/sqlc"
    "go-user-api/internal/service"
)

type UserHandler struct {
    db       *sqlc.Queries
    validate *validator.Validate
}

func NewUserHandler(db *sqlc.Queries) *UserHandler {
    return &UserHandler{
        db:       db,
        validate: validator.New(),
    }
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
    var req struct {
        Name string `json:"name" validate:"required"`
        DOB  string `json:"dob" validate:"required"`
    }

    if err := c.BodyParser(&req); err != nil {
        return fiber.ErrBadRequest
    }
    if err := h.validate.Struct(req); err != nil {
        return fiber.ErrBadRequest
    }

    user, err := h.db.CreateUser(c.Context(), sqlc.CreateUserParams{
        Name: req.Name,
        Dob:  req.DOB,
    })
    if err != nil {
        return fiber.ErrInternalServerError
    }
    return c.Status(201).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    user, err := h.db.GetUserByID(c.Context(), int32(id))
    if err != nil {
        return fiber.ErrNotFound
    }
    return c.JSON(fiber.Map{
        "id":   user.ID,
        "name": user.Name,
        "dob":  user.Dob.Format("2006-01-02"),
        "age":  service.CalculateAge(user.Dob),
    })
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
    users, _ := h.db.ListUsers(c.Context())
    res := []fiber.Map{}
    for _, u := range users {
        res = append(res, fiber.Map{
            "id":   u.ID,
            "name": u.Name,
            "dob":  u.Dob.Format("2006-01-02"),
            "age":  service.CalculateAge(u.Dob),
        })
    }
    return c.JSON(res)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    var req struct {
        Name string `json:"name"`
        DOB  string `json:"dob"`
    }
    if err := c.BodyParser(&req); err != nil {
        return fiber.ErrBadRequest
    }
    user, err := h.db.UpdateUser(c.Context(), sqlc.UpdateUserParams{
        ID:   int32(id),
        Name: req.Name,
        Dob:  req.DOB,
    })
    if err != nil {
        return fiber.ErrInternalServerError
    }
    return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
    id, _ := strconv.Atoi(c.Params("id"))
    _ = h.db.DeleteUser(c.Context(), int32(id))
    return c.SendStatus(204)
}
