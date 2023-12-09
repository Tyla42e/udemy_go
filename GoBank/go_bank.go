package main

import (
	"fmt"
)

func main() {
	var accountBalance = 1000.00
	var exit = false
	for !exit {
		var userChoice int
		displayMenu()
		userChoice = getUserChoice("Please select option: ")
		switch userChoice {
		case 1:
			fmt.Printf("Current Balance: $%.2f", accountBalance)
		case 2:
			{
				depositAmount := getAmount("Enter amount to deposit: $")
				accountBalance += depositAmount
				fmt.Printf("Current Balance: $%.2f", accountBalance)
			}
		case 3:
			{
				withdrawAmount := getAmount("Enter amount to withdraw: $")
				if withdrawAmount > accountBalance {
					fmt.Println("Not enough funds available")
				} else {
					accountBalance -= withdrawAmount
				}
				fmt.Printf("Current Balance: $%.2f", accountBalance)
			}
		case 4:
			exit = true
		default:
			fmt.Println("Invalid choice. Please enter a value between 1 - 4.")
		}
	}
	fmt.Println("Thank you for using the Bank of Go")
}

func getAmount(s string) float64 {
	var amount float64

	for {
		fmt.Print(s)

		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Invalid input: ", err)
		}
		if amount < 0 {
			fmt.Println("Please enter a positive value.")
		} else {
			break
		}
	}

	return amount
}

func displayMenu() {
	fmt.Println("\n\nWelcome to the Bank of Go!")
	fmt.Println("==========================")
	fmt.Println()
	fmt.Println("Menu")
	fmt.Println("====")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Make Deposit")
	fmt.Println("3. Make Withdrawal")
	fmt.Println("4. Quit")
	fmt.Println()
}

func getUserChoice(message string) int {
	var choice int
	fmt.Print("Please select option: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Input error: ", err)
		panic("")
	}
	return choice
}
