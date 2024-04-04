package main

import "fmt"
 
func main(){
	var accountBalance float64 = 1000



	fmt.Println("Welcome to Go Bank!")
	
	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice is: ")
		fmt.Scan(&choice)


		switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accountBalance)
		case 2:
			fmt.Println("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Deposit amount must be greater than 0")
				fmt.Println("Balance unchanged:", accountBalance)
				continue
			} 
			accountBalance += depositAmount
			fmt.Println("Balance updated! New balance:", accountBalance)
				
			case 3:
				fmt.Println("Withdraw amount: ")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)
				if withdrawAmount > accountBalance {
					fmt.Println("Insufficient funds for request")
					fmt.Println("Balance unchanged:", accountBalance)
					continue
				}
				accountBalance -= withdrawAmount
				fmt.Println("Balance updated! New balance:", accountBalance)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for banking with us today!")
				return
			}
		}	
}