package main

import (
	"fmt"
	"math"
)

func main() {
	// explicit declaration of variables
	// var <variable_name> <data_type> = <value>
	var investmentAmount float64 = 1000
	var years float64 = 10
	var expectedReturnRate = 5.5

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate/100, years) 

	fmt.Println(futureValue) // Println - Println adds a new line at the end of the output
}