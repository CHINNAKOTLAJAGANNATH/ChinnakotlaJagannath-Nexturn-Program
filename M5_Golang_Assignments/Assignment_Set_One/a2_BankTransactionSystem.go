/*
Exercise 2: Bank Transaction System
Topics Covered: Go Constants, Go Loops, Go Break and Continue, Go Functions, Go
Strings, Go Errors

Case Study:
You need to simulate a bank transaction system with the following features:
	1. Account Management: Each account has an ID, name, and balance. Store the
	accounts in a slice.

	2. Deposit Function: A function to deposit money into an account. Validate if the
	deposit amount is greater than zero.

	3. Withdraw Function: A function to withdraw money from an account. Ensure the
	account has a sufficient balance before proceeding. Return appropriate errors
	for invalid amounts or insufficient balance.

	4. Transaction History: Maintain a transaction history for each account as a string
	slice. Use a loop to display the transaction history when requested.

	5. Menu System: Implement a menu-driven program where users can choose
	actions like deposit, withdraw, view balance, or exit. Use constants for menu
	options and break the loop to exit.
*/

package main

import (
	"errors"
	"fmt"
)

type Account struct {
	ID              int
	Name            string
	Balance         float64
	TransactionHist []string
}

// Constants for menu options
const (
	DepositOption   = 1
	WithdrawOption  = 2
	ViewBalance     = 3
	TransactionHist = 4
	ExitOption      = 5
)

/* Deposit Function: A function to deposit money into an account. Validate if the
deposit amount is greater than zero. */
// Deposit money into the account
func Deposit(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Invalid deposit amount, deposit amount must be greater than zero.")
	}
	account.Balance += amount
	// show balancce
	fmt.Printf("Deposited $%.2f into account %d\n", amount, account.ID)
	fmt.Printf("Account Balance is: $%.2f\n", account.Balance)
	transaction := fmt.Sprintf("Deposited: $%.2f", amount)
	account.TransactionHist = append(account.TransactionHist, transaction)
	return nil
}

/* Withdraw Function: A function to withdraw money from an account. Ensure the
account has a sufficient balance before proceeding. Return appropriate errors
for invalid amounts or insufficient balance. */
// Withdraw money from the account
func Withdraw(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("Invalid withdrawal amount, withdrawal amount must be greater than zero.")
	}
	if account.Balance < amount {
		return errors.New("Insufficient balance.")
	}
	account.Balance -= amount
	fmt.Printf("Withdrawn $%.2f from account %d\n", amount, account.ID)
	fmt.Printf("Account Balance is: $%.2f\n", account.Balance)
	transaction := fmt.Sprintf("Withdrawn: $%.2f", amount)
	account.TransactionHist = append(account.TransactionHist, transaction)
	return nil
}

/* Transaction History: Maintain a transaction history for each account as a string
slice. Use a loop to display the transaction history when requested. */
// Display transaction history
func DisplayTransactionHistory(account *Account) {
	fmt.Println("Transaction History:")
	if len(account.TransactionHist) == 0 {
		fmt.Println("No transactions found.")
		return
	}
	for i, transaction := range account.TransactionHist {
		fmt.Printf("%d. %s\n", i+1, transaction)
	}
}

/* Menu System: Implement a menu-driven program where users can choose
actions like deposit, withdraw, view balance, or exit. Use constants for menu
options and break the loop to exit.  */
// Menu-driven program
func main() {
	accounts := []Account{
		{ID: 1, Name: "Jagan", Balance: 1000.00, TransactionHist: []string{}},
		{ID: 2, Name: "Hari", Balance: 500.00, TransactionHist: []string{}},
	}

	var accountID int
	var choice int

	fmt.Println("Welcome to the Bank Transaction System")
	for {
		fmt.Println("\nMenu:")
		fmt.Printf("%d. Deposit\n", DepositOption)
		fmt.Printf("%d. Withdraw\n", WithdrawOption)
		fmt.Printf("%d. View Balance\n", ViewBalance)
		fmt.Printf("%d. View Transaction History\n", TransactionHist)
		fmt.Printf("%d. Exit\n", ExitOption)

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == ExitOption {
			fmt.Println("Thank you for using the Bank Transaction System!")
			break
		}

		fmt.Print("Enter your account ID: ")
		fmt.Scan(&accountID)

		var account *Account
		for i := range accounts {
			if accounts[i].ID == accountID {
				account = &accounts[i]
				break
			}
		}

		if account == nil {
			fmt.Println("Invalid account ID.")
			continue
		}

		switch choice {
		case DepositOption:
			var amount float64
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			if err := Deposit(account, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful.")
			}
		case WithdrawOption:
			var amount float64
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			if err := Withdraw(account, amount); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful.")
			}
		case ViewBalance:
			fmt.Printf("Account Balance: $%.2f\n", account.Balance)
		case TransactionHist:
			DisplayTransactionHistory(account)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
