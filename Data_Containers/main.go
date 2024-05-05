package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"

	userNames = append(userNames, "Matt", "Max")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 10)

	courseRatings.output()

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 0.5
	
	courseRatings.output()

	// fmt.Println(courseRatings)

	for index, value := range userNames {}

}