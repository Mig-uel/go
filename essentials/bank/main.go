package main

import (
	"fmt"
	"os"
)

func main() {
	var userInput int
	var balance float64 = 1000

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")

		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		fmt.Print("Your choice: ")
		fmt.Scan(&userInput)

		switch userInput {
		case 1:
			fmt.Printf("Your balance is: $%.2f", balance)

		case 2:
			var depositAmount float64

			fmt.Print("Deposit amount: $")
			fmt.Scan(&depositAmount)

			if depositAmount < 1 {
				fmt.Println("Invalid deposit amount!")
				continue
			}

			balance += depositAmount

			fmt.Printf("Your new balance is: $%.2f\n", balance)
			writeToFile(balance)

		case 3:
			var withdrawAmount float64

			fmt.Print("Withdraw Amount: $")
			fmt.Scan(&withdrawAmount)

			if withdrawAmount > balance || withdrawAmount < 1 {
				fmt.Println("Invalid withdraw amount")
				continue
			}

			balance -= withdrawAmount

			fmt.Printf("Your new balance is: $%.2f\n", balance)
			writeToFile(balance)

		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing Go Bank!")

			return
		}
	}
}

func writeToFile(balance float64) {
	balanceString := fmt.Sprint(balance)

	os.WriteFile("balance.txt", []byte(balanceString), 0644)
}
