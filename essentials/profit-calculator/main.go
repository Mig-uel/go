package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64

	fmt.Println("*** Profit Calculator ***")

	// revenue
	printAndPass("Enter the revenue: ", &revenue)

	// expenses
	printAndPass("Enter total of expenses: ", &expenses)

	// tax rate
	printAndPass("Enter tax rate: ", &taxRate)

	ebt := revenue - expenses
	profit := ebt - (ebt * (taxRate / 100))
	ratio := ebt / profit

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}

func printAndPass(text string, x *float64) {
	fmt.Print(text)
	fmt.Scan(x)
}
