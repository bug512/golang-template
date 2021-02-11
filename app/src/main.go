package main

import (
	"fmt"
	"golang-template/codewars"
)

func main() {
	fmt.Println("It's ok!! This is template.")
	odd := codewars.FindOdd([]int{1, 2, 4, 1, 3, 3, 4, 5, 4, 4, 2, 6, 7, 7, 6})
	oddBest := codewars.FindOddBest([]int{1, 2, 4, 1, 3, 3, 4, 5, 4, 4, 2, 6, 7, 7, 6})

	fmt.Println(odd)
	fmt.Println(oddBest)
}
