package main

import (
	"fmt"
	"math"
)

func main() {
	// explicit declaration of variables
	// var <variable_name> <data_type> = <value>

	// var investmentAmount, years, expectedReturnRate float64 = 1000, 10, 5.5

	// var blocks can be used to group variables of the same type
	// var (
	// 	investmentAmount float64 = 1000
	// 	years int = 10
	// 	expectedReturnRate float64 = 5.5
	// )

	var investmentAmount float64 = 1000
	var years float64 = 10
	expectedReturnRate := 5.5 // type inference


	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years) 

	fmt.Println(futureValue) // Println - Println adds a new line at the end of the output
}