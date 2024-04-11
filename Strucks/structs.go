package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

func (u user) outputUserDetails(){
	// ...
	
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName(){
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name last name and birthdate required")
	}

	return &user{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate {MM/DD/YYYY}: ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()

	appUser.clearUserName()

	appUser.outputUserDetails()
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
