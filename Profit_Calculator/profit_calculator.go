package main

import "fmt"


func main(){
	var revenue, expenses, taxRate float64
	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter taxRate: ")
	fmt.Scan(&taxRate)

	var ebt float64 = revenue - expenses
	var profit float64 = ebt * (1 - taxRate / 100)
	var ratio float64 = ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)

}