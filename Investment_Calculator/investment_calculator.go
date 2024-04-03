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
	fmt.Print("Return: ")
	fmt.Println(futureValue)
	fmt.Print("Inflation Adjusted Return: ")
	fmt.Println(futureRealValue)
}