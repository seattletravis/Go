package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)
 
const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find file") 
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func writeFloatToFile(value float64, fileName string){
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func main(){
	var accountBalance, err = getFloatFromFile(accountBalanceFile)

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
			writeFloatToFile(accountBalance, accountBalanceFile)
				
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
				writeFloatToFile(accountBalance, accountBalanceFile)
			default:
				fmt.Println("Goodbye!")
				fmt.Println("Thanks for banking with us today!")
				return
			}
		}	
}

