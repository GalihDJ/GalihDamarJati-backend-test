package main

import "fmt"

func main() {

	input := 10

	var arrayNums []int

	for i := 1; i < input; i++ {
		result := sumNumbers(i)
		arrayNums = append(arrayNums, result)
	}

	sum := 0

	for i := 0; i <= input; i++ {
		if i != arrayNums[i] {
			sum += i
		}
	}

	fmt.Println(sum)

}

func sumNumbers(inputNumber int) int {

	output := inputNumber

	for inputNumber > 0 {
		output += inputNumber % 10
		inputNumber /= 10
	}

	return output
}
