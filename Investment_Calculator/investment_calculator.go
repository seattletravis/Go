package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64 = 1000
	var years float64 = 10.0
	var expectedReturnRate float64 = 5.5

	fmt.Print("Enter starting amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter interest rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter time in years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRealValue := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	// fmt.Println("Future return:", futureValue)
	// fmt.Printf("Future return: %.2f\nInflation Adjusted Return: %.2f\n", futureValue, futureRealValue)
	// fmt.Printf("Inflation Adjusted Return: %.0f\n", futureRealValue)

	fmt.Print(formattedFutureValue, formattedRealValue)

}
