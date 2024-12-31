package main

import (
	"a1_bms_project/config"
	"a1_bms_project/controller"
	"a1_bms_project/middleware"
	"a1_bms_project/repository"
	"a1_bms_project/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db := config.InitializeDatabase()

	// Set up repositories, services, and controllers
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	blogRepo := repository.NewBlogRepository(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	// Set up the router
	router := gin.Default()

	// Set up middleware for logging
	router.Use(middleware.LoggingMiddleware())

	// User routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	// Blog routes (require authentication)
	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware(db)) // Pass the database to AuthMiddleware
	{
		authorized.POST("/blog", middleware.ValidationMiddleware(), blogController.CreateBlog)
		authorized.GET("/blog/:id", blogController.GetBlog)
		authorized.GET("/blogs", blogController.GetAllBlogs)
		authorized.PUT("/blog/:id", middleware.ValidationMiddleware(), blogController.UpdateBlog)
		authorized.DELETE("/blog/:id", blogController.DeleteBlog)
	}

	// Run the server
	router.Run(":8080")
}
