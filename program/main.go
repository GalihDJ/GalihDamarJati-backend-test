package main

import "fmt"

func main() {

	input := 10

	var arrayNums []int

	for i := 1; i < input; i++ {
		result := sumNumbers(i)
		arrayNums = append(arrayNums, result)

		fmt.Println(arrayNums)
	}

	//sum := 0

}

func sumNumbers(inputNumber int) int {

	output := inputNumber

	for inputNumber > 0 {
		//fmt.Println(output)
		output += inputNumber % 10
		inputNumber /= 10
	}

	return output
}
