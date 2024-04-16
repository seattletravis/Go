package main

import (
	"fmt"
	"errors"
)

func main(){

}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Snanln(&value)

	if value == "" {
		return "", errors.New("invalid input")
	}

	return value, nil

}