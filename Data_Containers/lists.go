package main

import f "fmt"

func main(){
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99,9.99,45.99,20.0}
	f.Println(prices, productNames)

	productNames[2] = "A Carpet"

	f.Println(prices, productNames)

	featuredPrices := prices[1:3]

	f.Println(featuredPrices)
}