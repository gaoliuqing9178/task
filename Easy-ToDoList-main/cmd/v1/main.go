package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"todolist/config"
	"todolist/internal/database"
	"todolist/internal/handler"
	"todolist/internal/repository"
	"todolist/internal/router"
	"todolist/internal/service"
	"todolist/pkg/logUtil"
)

// @title			ToDo List API
// @version		1.0
// @description	This is a sample server for a todo list.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	fmt.Println("starting...")
	// 初始化配置文件
	config.Init()
	// 初始化日志
	logUtil.InitLogger()
	// 初始化数据库
	db := database.Init()
	logUtil.Logger.Info("配置文件、日志、数据库初始化完成")

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	toDoRepo := repository.NewToDoRepository(db)
	toDoService := service.NewToDoService(toDoRepo)
	toDoHandler := handler.NewToDoListHandler(toDoService)

	engine := gin.Default()

	// 设置路由
	router.SetupRouter(engine, userHandler, toDoHandler)

	// 从配置中获取服务端口
	serverPort := config.Get().Server.Port
	if serverPort == "" {
		serverPort = "8080" // 默认端口
	}

	logUtil.Logger.Info("服务启动于", zap.Any("port", serverPort))
	if err := engine.Run(":" + serverPort); err != nil {
		logUtil.Logger.Error("服务启动失败")
	}
}
