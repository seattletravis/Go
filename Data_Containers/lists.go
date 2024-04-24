package main

import f "fmt"

func main () {
	prices := []float64{10.99, 8.99}
	f.Println(prices[:1])

	prices = append(prices, 5.99, 4.99)

	f.Println(prices)

}


// func main(){
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99,9.99,45.99,20.0}
// 	f.Println(prices, productNames)

// 	productNames[2] = "A Carpet"

// 	f.Println(prices, productNames)

// 	featuredPrices := prices[:3]

// 	f.Println(featuredPrices)
// }