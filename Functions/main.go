package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 5, 4, 5, 6}

	summary := sumup(numbers)

	fmt.Println(summary)

}

func sumup(numbers []int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}