package main

import "fmt"

func main() {
	var revenue, expenses, taxRate float64
	var ebt, profit, ratio float64

	fmt.Println("*** Profit Calculator ***")

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter total of expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	ebt = revenue - expenses
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt / profit

	fmt.Println("EBT: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}
