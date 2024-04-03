package main

import "fmt"
 
func main(){
	var accountBalance float64 = 1000



	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice is: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your Balance is: ", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Println("Balance updated! New balance:", accountBalance)
	} else if choice == 3 {
		fmt.Println("Withdraw amount: ")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if (accountBalance - withdrawAmount) < 0 {
			fmt.Println("Insufficient funds for request")
			fmt.Println("Balance unchanged:", accountBalance)
		} else {
			accountBalance -= withdrawAmount
			fmt.Println("Balance updated! New balance:", accountBalance)
		}
	} else {
		fmt.Println("Thanks for banking with us today!")
	}
}