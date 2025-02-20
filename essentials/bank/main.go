package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const balanceFile = "balance.txt"

func main() {
	var userInput int
	var balance, err = readFromFile()

	if err != nil {
		fmt.Print("ERROR: ")
		fmt.Println(err)
		fmt.Println("-----------")
	}

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
			fmt.Printf("Your balance is: $%.2f\n", balance)

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

// wite to file
func writeToFile(balance float64) {
	balanceString := fmt.Sprint(balance)

	os.WriteFile(balanceFile, []byte(balanceString), 0644)
}

// read from file
func readFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}
