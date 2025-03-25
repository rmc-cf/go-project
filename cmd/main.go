package main

import (
    "myapi/cmd/api/routes"
    "myapi/internal/repository"
    "myapi/internal/service"
    "myapi/pkg/config"
    "myapi/pkg/logger"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
    // 初始化配置和日志
    config.LoadConfig("configs/config.yaml")
    log := logger.New()

    // 初始化数据库
    db := initDatabase()

    // 创建Fiber应用
    app := fiber.New(fiber.Config{
        ErrorHandler: globalErrorHandler, // 全局错误处理
    })
    app.Use(logger.New()) // Fiber内置日志中间件

    // 依赖注入 & 注册路由
    registerRoutes(app, db)

    // 启动服务
    startServer(app)
}

// 初始化数据库连接
func initDatabase() *gorm.DB {
    dsn := config.GetString("database.dsn")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        logger.New().Fatal("数据库连接失败", zap.Error(err))
    }
    return db
}

// 注册所有路由
func registerRoutes(app *fiber.App, db *gorm.DB) {
    // 用户模块
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    routes.RegisterUserRoutes(app, userService)

    // 其他实体路由可按相同模式添加：
    // productRepo := repository.NewProductRepository(db)
    // productService := service.NewProductService(productRepo)
    // routes.RegisterProductRoutes(app, productService)
}

// 全局错误处理器
func globalErrorHandler(c *fiber.Ctx, err error) error {
    code := fiber.StatusInternalServerError
    if e, ok := err.(*fiber.Error); ok {
        code = e.Code
    }
    return c.Status(code).JSON(fiber.Map{
        "error": err.Error(),
    })
}

// 启动HTTP服务
func startServer(app *fiber.App) {
    port := config.GetString("server.port")
    if err := app.Listen(":" + port); err != nil {
        logger.New().Fatal("服务启动失败", zap.Error(err))
    }
}