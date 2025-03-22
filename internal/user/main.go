package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jiahuipaung/Codefolio_Backend/common/config"
	"github.com/jiahuipaung/Codefolio_Backend/user/adapters/database"
	"github.com/jiahuipaung/Codefolio_Backend/user/adapters/memory"
	"github.com/jiahuipaung/Codefolio_Backend/user/app"
	"github.com/jiahuipaung/Codefolio_Backend/user/domain"
	"github.com/jiahuipaung/Codefolio_Backend/user/ports"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}
	cfg := config.GetConfig()

	// 根据配置选择存储方式
	var userRepo domain.UserRepository
	switch cfg.Storage.Type {
	case "database":
		// 初始化数据库连接
		db, err := database.NewDB(cfg.Storage.Database.GetDatabaseDSN())
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		userRepo = database.NewUserRepository(db)
	case "memory":
		fallthrough
	default:
		// 使用内存存储
		userRepo = memory.NewUserRepository()
	}

	// 初始化服务
	userService := app.NewUserService(userRepo, cfg.Server.JWTSecret)

	// 初始化 HTTP 处理器
	handler := ports.NewHandler(userService)

	// 配置路由
	r := gin.Default()

	// 用户相关路由
	userGroup := r.Group("/api/v1/users")
	{
		userGroup.POST("/register", handler.Register)
		userGroup.POST("/login", handler.Login)
	}

	// 启动服务器
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
