package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
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

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name last name and birthdate required")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}