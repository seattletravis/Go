package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64 = 1000
	var years float64 = 10.0
	var expectedReturnRate float64 = 5.5

	fmt.Print("Enter starting amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter interest rate: ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter time in years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFutureValue := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRealValue := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	// fmt.Println("Future return:", futureValue)
	// fmt.Printf("Future return: %.2f\nInflation Adjusted Return: %.2f\n", futureValue, futureRealValue)
	// fmt.Printf("Inflation Adjusted Return: %.0f\n", futureRealValue)

	// fmt.Print(formattedFutureValue, formattedRealValue)

	outputText(formattedFutureValue, formattedRealValue)

}

func outputText(text1, text2 string){
	fmt.Print(text1, text2)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64){
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years) 
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}