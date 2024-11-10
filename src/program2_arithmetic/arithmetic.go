package arithmetic /*A program of Basic Arithmetic Operations in Go*/

import (
	"fmt"
	"math"
)

func ProgramTwoArithmetic() {
	fmt.Println("Go Arithmetic")

	fmt.Println("This program runs basic arithmetic operations like +,*,/,-,** ")

	// Addition
	sum := add(50, 50)

	// Subtraction
	sub(2, 2)

	// Division
	div(100, 100)

	// Multiplication
	mul(5, 5)

	// Exponentiation
	res := pow(2.1, 2.1)
	fmt.Printf(" 2.1 ** 2.1 = %v ", res)

	fmt.Println(sum)
}

func add(num1 int8, num2 int8) (sum int8) {

	fmt.Println("Running Addition")

	var result int8 = num1 + num2
	fmt.Printf(" %v + %v = %v \n", num1, num2, result)
	return result
}

func sub(num1 int8, num2 int8) (diff int8) {

	fmt.Println("Running Subtraction")

	var result int8 = num1 - num2
	fmt.Printf(" %v - %v = %v \n", num1, num2, result)
	return result
}

func div(num1 int8, num2 int8) (sum int8) {

	fmt.Println("Running Division")

	var result int8 = num1 / num2
	fmt.Printf(" %v / %v = %v \n", num1, num2, result)
	return result
}

func mul(num1 int8, num2 int8) (sum int8) {

	fmt.Println("Running Multiplication")

	var result int8 = num1 * num2
	fmt.Printf(" %v * %v = %v \n", num1, num2, result)
	return result
}

func pow(num float64, power float64) float64 {

	fmt.Println("Running Exponentiation")
	return math.Pow(num, power)
}
