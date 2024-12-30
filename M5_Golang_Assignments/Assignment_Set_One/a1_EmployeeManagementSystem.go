/*
Employee Management System

Topics Covered: Go Conditions, Go Loops, Go Constants, Go Functions, Go Arrays, Go
Strings, Go Errors

Case Study:
	A company wants to manage its employees' data in memory. Each employee has an ID,
	name, age, and department. You need to build a small application that performs the
	following:
	1. Add Employee: Accept input for employee details and store them in an array of
	structs. Validate the input:
		o ID must be unique.
		o Age should be greater than 18. If validation fails, return custom error
		messages.
	2. Search Employee: Search for an employee by ID or name using conditions.
	Return the details if found, or return an error if not found.
	3. List Employees by Department: Use loops to filter and display all employees in
	a given department.
	4. Count Employees: Use constants to define a department (e.g., "HR", "IT"), and
	display the count of employees in that department.

Bonus:
Refactor the repetitive code using functions, and add error handling for invalid
operations like searching for a non-existent employee.

*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Employee struct to hold employee details
type Employee struct {
	ID         string
	Name       string
	Age        int
	Department string
}

// Constants for departments
const (
	HR    = "HR"
	IT    = "IT"
	Sales = "Sales"
)

// Global slice to store employees
var employees []Employee

/* 1. Add Employee: Accept input for employee details and store them in an array of
structs. Validate the input:
	o ID must be unique.
	o Age should be greater than 18. If validation fails, return custom error
	messages.
*/
// AddEmployee adds a new employee after validation
func addEmployee(id string, name string, age int, department string) error {
	// Check if ID already exists
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("Employee ID already exists, employee ID must be unique")
		}
	}
	// Validate age
	if age <= 18 {
		return errors.New("Age should be greater than 18")
	}

	// Create a new employee and add to the slice
	newEmployee := Employee{ID: id, Name: name, Age: age, Department: department}
	employees = append(employees, newEmployee)
	return nil
}

/* 2. Search Employee: Search for an employee by ID or name using conditions.
Return the details if found, or return an error if not found.
*/
// SearchEmployee searches for an employee by ID or name
func searchEmployee(idOrName string) (*Employee, error) {
	// Iterate through the employees slice to find the match
	for _, emp := range employees {
		if emp.ID == idOrName || emp.Name == idOrName {
			return &emp, nil
		}
	}
	return nil, errors.New("Employee not found")
}

/* 3. List Employees by Department: Use loops to filter and display all employees in
a given department. */
// ListEmployeesByDepartment lists employees by department
func listEmployeesByDepartment(department string) []Employee {
	var filteredEmployees []Employee
	// Loop through the employees slice to find the match
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			// fmt.Printf("ID: %s, Name: %s, Age: %d, Department
			// ", emp.ID, emp.Name, emp.Age)
			filteredEmployees = append(filteredEmployees, emp)
		}
	}
	return filteredEmployees
}

/* 4. Count Employees: Use constants to define a department (e.g., "HR", "IT"), and
display the count of employees in that department. */
// CountEmployees counts employees by department
func countEmployees(department string) int {
	// Initialize a counter variable
	count := 0
	// Loop through the employees slice to count the match
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

// DisplayEmployee prints employee details
func displayEmployee(emp Employee) {
	fmt.Printf("ID: %s, Name: %s, Age: %d, Department: %s\n", emp.ID, emp.Name, emp.Age, emp.Department)
}

// CaptureInput captures user input for an employee
func captureInput() (string, string, int, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Employee ID: ")
	id, _ := reader.ReadString('\n')
	id = strings.TrimSpace(id)

	fmt.Print("Enter Employee Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Employee Age: ")
	ageStr, _ := reader.ReadString('\n')
	ageStr = strings.TrimSpace(ageStr)
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return "", "", 0, "", errors.New("invalid age format")
	}

	fmt.Print("Enter Employee Department: ")
	department, _ := reader.ReadString('\n')
	department = strings.TrimSpace(department)

	return id, name, age, department, nil
}

// Main function to test the system
func main() {
	for {
		fmt.Println("\nEmployee Management System")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. List Employees by Department")
		fmt.Println("4. Count Employees in Department")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			// Add Employee
			id, name, age, department, err := captureInput()
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			err = addEmployee(id, name, age, department)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee added successfully!")
			}

		case 2:
			// Search Employee
			fmt.Print("Enter Employee ID or Name to search: ")
			var query string
			fmt.Scanln(&query)
			emp, err := searchEmployee(query)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee found:")
				displayEmployee(*emp)
			}

		case 3:
			// List Employees by Department
			fmt.Print("Enter Department to list employees: ")
			var department string
			fmt.Scanln(&department)
			employees := listEmployeesByDepartment(department)
			if len(employees) == 0 {
				fmt.Println("No employees found in the department.")
			} else {
				fmt.Println("Employees in", department, "department:")
				for _, emp := range employees {
					displayEmployee(emp)
				}
			}

		case 4:
			// Count Employees
			fmt.Print("Enter Department to count employees: ")
			var department string
			fmt.Scanln(&department)
			count := countEmployees(department)
			fmt.Printf("Number of employees in %s department: %d\n", department, count)

		case 5:
			// Exit
			fmt.Println("Exiting the system. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
