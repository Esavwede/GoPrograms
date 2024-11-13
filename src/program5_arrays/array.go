package arrays

import "fmt"

func Program5Arrays() {
	fmt.Println("Arrays")

	var myNumbers = [...]int8{1, 2, 3}

	fmt.Printf("This are my numbers: %v \n", myNumbers)

	fmt.Printf("The first element in my array is: %v \n", myNumbers[0])
}
