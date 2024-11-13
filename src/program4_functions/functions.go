package functions

import (
	"errors"
	"fmt"
)

func Program4Functions() {
	fmt.Println("Functions")

	// Division
	var numerator int = 10
	var denominator int = 0

	dividend, e := divider(numerator, denominator)

	if e != nil {
		fmt.Printf("%v \n", e.Error())
	} else {
		fmt.Printf("%v / %v = %v \n", numerator, denominator, dividend)
	}
}

func divider(numerator int, denominator int) (int, error) {

	if denominator == 0 {
		e := errors.New("zero division error ")

		return 0, e
	}

	return numerator / denominator, nil
}
