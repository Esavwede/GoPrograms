package conditionals

import "fmt"

func Program3Conditionals() {
	cond1()
	cond2()
	cond3()
}

func cond1() {
	var specialNumber int8 = 101

	if specialNumber < 100 {
		fmt.Printf("%v is less than a 100 \n", specialNumber)
	} else {
		fmt.Printf("%v is greater than a 100 \n", specialNumber)
	}
}

func cond2() {
	// Number range
	var specialLetter string = "c"

	if specialLetter == "a" {
		fmt.Printf("Special letter is a \n")
	} else if specialLetter == "b" {
		fmt.Printf("Special letter is b \n")
	} else if specialLetter == "c" {
		fmt.Printf("Special letter is c \n")
	} else {
		fmt.Printf("Unknown Special Character \n")
	}
}

func cond3() {

	myNum := 2
	myChar := "z"
	userSignedUp := true

	if myNum < 100 && myChar == "o" && userSignedUp {
		fmt.Printf("Correct character combination \n")
	} else {
		fmt.Printf("Incorrect character combination \n")
	}
}
