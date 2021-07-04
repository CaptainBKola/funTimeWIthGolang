package main

import "fmt"

func main() {
	fmt.Println("hello there, this is my first algorithm practice in Go")

	fmt.Print(findListOfNumbers([]int{1, 234, 5, 7, 5}, 6))
}
func findListOfNumbers(lis []int, number int) bool {
	for _, value := range lis {

		if value == number {

			return true

		}
	}

	return false

}
