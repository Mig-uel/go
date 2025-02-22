package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue, expenses, taxRate float64

	fmt.Println("*** Profit Calculator ***")

	// revenue
	printAndPass("Enter the revenue: ", &revenue)
	err := validate(revenue)
	if err != nil {
		panic(err)
	}

	// expenses
	printAndPass("Enter total of expenses: ", &expenses)
	err = validate(expenses)
	if err != nil {
		panic(err)
	}

	// tax rate
	printAndPass("Enter tax rate: ", &taxRate)
	err = validate(expenses)
	if err != nil {
		panic(err)
	}

	// calculate
	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.3f", ratio)
}

func validate(input float64) error {
	if input <= 0 {
		return errors.New("input cannot be less than or equal to 0")
	}

	return nil
}

func printAndPass(text string, x *float64) {
	fmt.Print(text)
	fmt.Scan(x)
}

func calculate(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64) {

	ebt = revenue - expenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt / profit

	store(ebt, profit, ratio)

	return ebt, profit, ratio
}

func store(ebt, profit, ratio float64) {
	s := fmt.Sprintf("EBT: %.2f PROFIT: %.2f RATIO: %.3f", ebt, profit, ratio)

	err := os.WriteFile("results.txt", []byte(s), 0664)

	if err != nil {
		panic(err)
	}
}
