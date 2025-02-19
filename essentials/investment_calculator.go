package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount, years = 1000, 10
	var expectedReturnRate = 5.5

	var futureValue = float64(investmentAmount) * math.Pow((1 + expectedReturnRate/float64(years)), expectedReturnRate) 

	fmt.Println(futureValue) // Println - Println adds a new line at the end of the output
}