package main

import "fmt"


func main(){
	var revenue, expenses, taxRate float64 = 1, 1, 1
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter taxRate: ")
	fmt.Scan(&taxRate)

	var profit float64 = revenue - revenue * taxRate / 100 - expenses

	fmt.Print(profit)
	
}