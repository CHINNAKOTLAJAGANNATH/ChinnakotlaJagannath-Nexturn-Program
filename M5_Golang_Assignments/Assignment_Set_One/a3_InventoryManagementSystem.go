/* Exercise 3: Inventory Management System
Topics Covered: Go Conditions, Go Type Casting, Go Functions, Go Arrays, Go Strings,
Go Errors

Case Study:
A store needs to manage its inventory of products. Build an application that includes
the following:
	1. Product Struct: Create a struct to represent a product with fields for ID, name,
	price (float64), and stock (int).

	2. Add Product: Write a function to add new products to the inventory. Use type
	casting to ensure price inputs are converted to float64.

	3. Update Stock: Implement a function to update the stock of a product. Use
	conditions to validate the input (e.g., stock cannot be negative).

	4. Search Product: Allow users to search for products by name or ID. If a product is
	not found, return a custom error message.

	5. Display Inventory: Use loops to display all available products in a formatted
	table.
Bonus:
• Add sorting functionality to display products by price or stock in ascending order. */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* Product Struct: Create a struct to represent a product with fields for ID, name,
price (float64), and stock (int). */
// Product struct represents a product in the inventory
type Product struct {
	ID    string
	Name  string
	Price float64
	Stock int
}

// Inventory holds all the products
var Inventory []Product

/* Add Product: Write a function to add new products to the inventory. Use type
casting to ensure price inputs are converted to float64.*/
// AddProduct adds a new product to the inventory
func addProduct(id, name, priceInput string, stock int) error {
	// Convert price to float64 using type casting
	priceFloat, err := strconv.ParseFloat(priceInput, 64)
	if err != nil {
		return fmt.Errorf("Invalid price: %v", err)
	}
	if stock < 0 {
		return errors.New("stock cannot be negative")
	}
	// Create a new product and add it to the inventory
	newProduct := Product{ID: id, Name: name, Price: priceFloat, Stock: stock}
	Inventory = append(Inventory, newProduct)
	return nil
}

/* Update Stock: Implement a function to update the stock of a product. Use
conditions to validate the input (e.g., stock cannot be negative). */
// UpdateStock updates the stock of a product
func updateStock(id string, newStock int) error {
	if newStock < 0 {
		return errors.New("stock cannot be negative")
	}
	for i := range Inventory {
		if Inventory[i].ID == id {
			// Update the stock of the product
			Inventory[i].Stock = newStock
			return nil
		}
	}
	return fmt.Errorf("product with ID %s not found", id)
}

/* Search Product: Allow users to search for products by name or ID. If a product is
not found, return a custom error message. */
// SearchProduct searches for a product by name or ID
func searchProduct(idOrName string) (*Product, error) {
	for _, product := range Inventory {
		if product.ID == idOrName || product.Name == idOrName {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product with ID or name %s not found", idOrName)
}

/* Display Inventory: Use loops to display all available products in a formatted
table. */
//  DisplayInventory displays all products in a formatted table
func displayInventory() {
	fmt.Printf("%-10s %-20s %-10s %-10s\n", "ID", "Name", "Price", "Stock")
	fmt.Println(strings.Repeat("-", 50))
	for _, product := range Inventory {
		fmt.Printf("%-10s %-20s %-10.2f %-10d\n", product.ID, product.Name, product.Price, product.Stock)
	}
}

// sortInventory sorts the inventory by price or stock in ascending order
func sortInventory(by string) {
	switch by {
	case "price":
		sort.Slice(Inventory, func(i, j int) bool {
			return Inventory[i].Price < Inventory[j].Price
		})
	case "stock":
		sort.Slice(Inventory, func(i, j int) bool {
			return Inventory[i].Stock < Inventory[j].Stock
		})
	default:
		fmt.Println("Invalid sort option. Use 'price' or 'stock'.")
	}
}

// Get user input
func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
func main() {
	for {
		fmt.Println("\nInventory Management System")
		fmt.Println("1. Add Product")
		fmt.Println("2. Update Stock")
		fmt.Println("3. Search Product")
		fmt.Println("4. Display Inventory")
		fmt.Println("5. Sort Inventory")
		fmt.Println("6. Exit")
		choice := getUserInput("Enter your choice: ")

		switch choice {
		case "1":
			id := getUserInput("Enter Product ID: ")
			name := getUserInput("Enter Product Name: ")
			price := getUserInput("Enter Product Price: ")
			stockInput := getUserInput("Enter Product Stock: ")
			stock, err := strconv.Atoi(stockInput)
			if err != nil {
				fmt.Println("Invalid stock value.")
				break
			}
			err = addProduct(id, name, price, stock)
			if err != nil {
				fmt.Println("Error adding product:", err)
			} else {
				fmt.Println("Product added successfully.")
			}
		case "2":
			id := getUserInput("Enter Product ID to update: ")
			stockInput := getUserInput("Enter new stock value: ")
			stock, err := strconv.Atoi(stockInput)
			if err != nil {
				fmt.Println("Invalid stock value.")
				break
			}
			err = updateStock(id, stock)
			if err != nil {
				fmt.Println("Error updating stock:", err)
			} else {
				fmt.Println("Stock updated successfully.")
			}
		case "3":
			query := getUserInput("Enter Product ID or Name to search: ")
			product, err := searchProduct(query)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Product Found: %+v\n", *product)
			}
		case "4":
			fmt.Println("\nInventory:")
			displayInventory()
		case "5":
			sortBy := getUserInput("Sort by 'price' or 'stock': ")
			sortInventory(sortBy)
			fmt.Println("Inventory sorted successfully.")
			displayInventory()
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
