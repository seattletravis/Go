package main

import (
	"fmt"

	"example.com/structs/user"
)



func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate {MM/DD/YYYY}: ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()

	appUser.ClearUserName()

	appUser.OutputUserDetails()
}

// func outputUserDetails(u *user){
// 	fmt.Println(u.firstName, u.lastName, u.birthDate)
// }

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
