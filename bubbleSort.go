package main

import "fmt"

var numbers = []int{5, 3, 2, 5, 4, 2, 1}

var k int

func main() {
	bubbleSort(numbers)
}

func bubbleSort(numbers []int) {
	var N = len(numbers)

	for i := 0; i < N; i++ {
		if !sweep(numbers, i) {
			return
		}

		fmt.Println(numbers)
	}
}

func sweep(numbers []int, prevPasses int) bool {
	var (
		N           = len(numbers)
		firstIndex  = 0
		secondIndex = 1
		didSwap     = false
	)
	for secondIndex < (N - prevPasses) {
		var firstNumber = numbers[firstIndex]
		var secondNumber = numbers[secondIndex]

		//do swap
		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
			didSwap = true
		}
		firstIndex++
		secondIndex++
	}
	return didSwap
}
