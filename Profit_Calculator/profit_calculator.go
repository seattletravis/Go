package main

import (
	"errors"
	"fmt"
)


func main(){

	revenue, err1 := promptUser("Enter revenue: ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	expenses, err2 := promptUser("Enter expenses: ")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	taxRate, err3 := promptUser("Enter tax rate: ")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	ebt, profit, ratio := calculate(revenue, expenses, taxRate)

	fmt.Printf("$%.2f\n",ebt)
	fmt.Printf("$%.2f\n",profit)
	fmt.Printf("%.2f\n",ratio)
}

func promptUser(text string)(float64, error){
	var output float64
	fmt.Print(text)
	fmt.Scan(&output)
	if output <= 0 {
		return 0, errors.New("value must be a positive number")
	}
	return output, nil
}

func calculate(revenue, expenses, taxRate float64) (ebt float64, profit float64, ratio float64){
	ebt = revenue - expenses
	profit = ebt * (1  -taxRate / 100)
	ratio = ebt / profit
	return ebt, profit, ratio
}