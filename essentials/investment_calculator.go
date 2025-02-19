package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Enter the investment amount: $")
	fmt.Scan(&investmentAmount)

	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedFutureRealValue := fmt.Sprintf("Realistic Future Value: %.2f\n", futureRealValue)

	fmt.Print(formattedFutureValue)
	fmt.Print(formattedFutureRealValue)
}

func outputText(input string) {
	fmt.Print(input)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	frv := fv / math.Pow(1+inflationRate/100, years)

	return fv, frv
}
