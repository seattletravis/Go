package main

import "fmt"


func main(){

	revenue := promptUser("Enter revenue: ")
	expenses := promptUser("Enter expenses: ")
	taxRate := promptUser("Enter tax rate: ")

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("$%.2f\n",ebt)
	fmt.Printf("$%.2f\n",profit)
	fmt.Printf("%.2f\n",ratio)

}

func promptUser(text string)(float64){
	var output float64
	fmt.Print(text)
	fmt.Scan(&output)
	return output
}

func calculate(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1  -taxRate / 100)
	ratio = ebt / profit
	return ebt, profit, ratio
}