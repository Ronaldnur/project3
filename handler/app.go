package handler

import (
	"fmt"
	"project3/infra/config"
	"project3/infra/database"
	"project3/repository/category_repository/category_pg"
	"project3/repository/task_repository/task_pg"
	"project3/repository/user_repository/user_pg"
	"project3/service"

	"github.com/gin-gonic/gin"
)

func SeedAdmin() {
	config.LoadAppConfig()

	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	_, err := userService.SeedAdminUser()
	if err != nil {
		fmt.Printf("Error seeding admin user: %v\n", err)
		return
	}
}
func StartApp() {
	config.LoadAppConfig()

	database.InitiliazeDatabase()

	var port = config.GetAppConfig().Port

	db := database.GetDatabaseInstance()

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	categoryRepo := category_pg.NewCategoryPG(db)

	categoryService := service.NewCategoryService(categoryRepo)

	categoryHandler := NewCategoryHandler(categoryService)

	taskRepo := task_pg.NewTaskPG(db)

	taskService := service.NewTaskService(taskRepo)

	taskHandler := NewTaskHandler(taskService)

	authService := service.NewAuthService(userRepo, categoryRepo, taskRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("login", userHandler.Login)
		userRoute.Use(authService.Authentitaction())
		userRoute.PUT("/update-account", userHandler.UpdateUser)
		userRoute.PATCH("/", userHandler.PatchRole)
		userRoute.DELETE("/delete-account", userHandler.DeleteUser)
	}

	categoryRoute := route.Group("/categories")
	{
		categoryRoute.Use(authService.Authentitaction())
		categoryRoute.POST("/", authService.AuthorizationAdmin(), categoryHandler.CreateCategory)
		categoryRoute.GET("/", categoryHandler.GetCategory)
		categoryRoute.PATCH("/:categoryId", authService.AuthorizationAdmin(), categoryHandler.PatchCategory)
		categoryRoute.DELETE("/:categoryId", authService.AuthorizationAdmin(), categoryHandler.DeleteCategory)
	}

	taskRoute := route.Group("/tasks")
	{
		taskRoute.Use(authService.Authentitaction())
		taskRoute.POST("/", taskHandler.CreateTask)
		taskRoute.PUT("/:taskId", authService.AuthorizationTask(), taskHandler.UpdateTask)
		taskRoute.PATCH("/update-status/:taskId", authService.AuthorizationTask(), taskHandler.PatchStatus)
		taskRoute.PATCH("/update-category/:taskId", authService.AuthorizationTask(), taskHandler.PatchCategoryId)
		taskRoute.DELETE(":taskId", authService.AuthorizationTask(), taskHandler.DeleteTask)
		taskRoute.GET("/", taskHandler.GetTask)
	}

	route.Run(":" + port)

}
