package main

import (
	"a2_ems_project/config"
	"a2_ems_project/controller"
	"a2_ems_project/middleware"
	"a2_ems_project/repository"
	"a2_ems_project/service"
	"net/http"
)

func main() {
	// Initialize database
	db := config.InitializeDatabase()

	// Set up repositories, services, and controllers
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	// Set up router
	mux := http.NewServeMux()

	// User routes
	mux.HandleFunc("/register", userController.Register)
	mux.HandleFunc("/login", userController.Login)

	// Product routes
	productHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			if r.URL.Path == "/product" {
				productController.AddProduct(w, r)
			} else {
				http.NotFound(w, r)
			}
		case http.MethodGet:
			if r.URL.Path == "/products" {
				productController.GetAllProducts(w, r)
			} else if len(r.URL.Path) > 9 && r.URL.Path[:9] == "/product/" {
				productController.GetProduct(w, r)
			} else {
				http.NotFound(w, r)
			}
		case http.MethodPut:
			if len(r.URL.Path) > 9 && r.URL.Path[:9] == "/product/" {
				productController.UpdateProduct(w, r)
			} else {
				http.NotFound(w, r)
			}
		case http.MethodDelete:
			if len(r.URL.Path) > 9 && r.URL.Path[:9] == "/product/" {
				productController.DeleteProduct(w, r)
			} else {
				http.NotFound(w, r)
			}
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Middleware for products
	securedProductHandler := middleware.LoggingMiddleware(
		middleware.ValidationMiddleware(
			middleware.AuthMiddleware(productHandler),
		),
	)

	// Register routes
	mux.Handle("/product", securedProductHandler)
	mux.Handle("/products", securedProductHandler)
	mux.Handle("/product/", securedProductHandler)

	// Start the server on port 8080
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
