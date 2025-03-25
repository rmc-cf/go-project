package routes

import (
    "myapi/internal/handler"
    "myapi/internal/service"

    "github.com/gofiber/fiber/v2"
)

// RegisterUserRoutes 用户实体的路由注册
func RegisterUserRoutes(app *fiber.App, userService service.UserService) {
    // 初始化Handler
    userHandler := handler.NewUserHandler(userService)

    // 路由分组
    api := app.Group("/api/v1")
    {
        api.Post("/users", userHandler.CreateUser)
        api.Get("/users/:id", userHandler.GetUser)
    }
}
