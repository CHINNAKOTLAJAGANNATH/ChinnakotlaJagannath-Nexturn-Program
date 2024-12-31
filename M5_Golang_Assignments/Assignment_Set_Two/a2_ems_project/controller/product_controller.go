package controller

import (
	"a2_ems_project/model"
	"a2_ems_project/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ProductController struct {
	ProductService *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{ProductService: service}
}

func (controller *ProductController) AddProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := controller.ProductService.AddProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product added successfully"})
}

func (controller *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	// Extract product ID from the URL path
	id := r.URL.Path[len("/product/"):]

	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := controller.ProductService.GetProductByID(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Set Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Indent the JSON response with 4 spaces for pretty printing
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	encoder.Encode(product)
}

func (controller *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL path using the URL structure "/product/{id}"
	id := r.URL.Path[len("/product/"):]

	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var product model.Product
	product.ID = productID
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = controller.ProductService.UpdateProduct(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}

func (controller *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL path using the URL structure "/product/{id}"
	id := r.URL.Path[len("/product/"):]

	// Convert the product ID to integer
	productID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Call the service method to delete the product
	err = controller.ProductService.DeleteProduct(productID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}

func (controller *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	// Get all products from the service layer
	products, err := controller.ProductService.GetAllProducts(page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the response header for JSON content
	w.Header().Set("Content-Type", "application/json")

	// Loop through the products and print each on a new line
	for _, product := range products {
		jsonData, err := json.Marshal(product)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Print each product on a new line
		fmt.Fprintln(w, string(jsonData))
	}
}
