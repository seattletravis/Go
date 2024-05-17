package main

import (
	"fmt"

	"example.com/price-calculator/prices"
) 

func main() {
	taxRate := []float64{0, 0.7, 0.1, 0.15}

	result := make(map[float64][]float64)

	for _, taxRate := range taxRate {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	fmt.Println(result)
}