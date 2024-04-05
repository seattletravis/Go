package main

import (
	"fmt"

	"example.com/bank/fileops"
)
 
const accountBalanceFile = "balance.txt"

func main(){
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------------")
	}

	fmt.Println("Welcome to Go Bank!")
	
	for {

		presentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
				
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for banking with us today!")
				return
			}
		}	
}

