package main

import "fmt"

// Sum calculates the total from an array of numbers.
func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {
	numbers := [5]int{1, 2, 4, 7, 50}
	fmt.Println(Sum(numbers))
	fmt.Println(numbers)

}
