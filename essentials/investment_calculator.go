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

	// fmt.Print("Enter the investment amount: $")
	outputText("Enter the investment amount: $")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter the number of years: ")
	outputText("Enter the number of years: ")
	fmt.Scan(&years)

	// fmt.Print("Enter the expected return rate: ")
	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// futureValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %.2f\n", futureValue)
	// fmt.Printf("Realistic Future Value: %.2f\n", futureRealValue)

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
