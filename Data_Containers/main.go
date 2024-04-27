package main

import "fmt"

func main() {
	userNames := make([]string, 2)

	userNames[0] = "Juwel"

	userNames = append(userNames, "Matt", "Max")

	fmt.Println(userNames)
}