package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "todolist/docs" // swagger docs
	"todolist/internal/handler"
	"todolist/internal/middleware"
)

// SetupRouter configures the routing for the application.
func SetupRouter(engine *gin.Engine, userHandler *handler.UserHandler, toDoHandler *handler.ToDoListHandler) {
	// Swagger endpoint
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API v1 group
	v1 := engine.Group("/api/v1")
	{
		// Public routes for user registration and login
		v1.POST("/user/register", userHandler.Register)
		v1.POST("/user/login", userHandler.Login)

		// Authenticated routes
		authed := v1.Group("/")
		authed.Use(middleware.JWTAuthMiddleware())
		{
			// To-Do routes
			authed.GET("/todos/list", toDoHandler.ListByUser)
			authed.POST("/todos/add", toDoHandler.Create)
			authed.GET("/todos/get/:id", toDoHandler.GetByID)
			authed.PUT("/todos/update:id", toDoHandler.UpdateByID)
			authed.DELETE("/todos/delete/:id", toDoHandler.Delete)

			// User profile update
			authed.PUT("/user/update", userHandler.UpdateUserInfo)
		}
	}
}
