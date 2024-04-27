package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Juwel"

	userNames = append(userNames, "Matt", "Max")

	fmt.Println(userNames)

	courseRatings := map[string]float64{}

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8

}