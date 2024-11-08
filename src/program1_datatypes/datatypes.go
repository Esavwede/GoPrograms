package program1

import "fmt"

func Datatypes() {
	fmt.Println("Basic Datatypes")
	integers()
	floats()
	bools()
	strings()
}

func integers() {

	fmt.Println("--Integers--")

	// signed Integers
	var myInt8 int8 = 127
	var myInt16 int16 = 10000
	var myInt32 int32 = 1000000000
	var myInt64 int64 = 1000000000000000000

	// Unsigned Integers
	var myUInt8 uint8 = 100
	var myUInt16 uint16 = 10000
	var myUInt32 uint32 = 100000000
	var myUInt64 uint64 = 10000000000000

	// Signed Integers
	fmt.Printf("My 8 bit integer is: %d \n", myInt8)
	fmt.Printf("My 16 bit integer is: %d \n", myInt16)
	fmt.Printf("My 32 bit integer is: %d \n", myInt32)
	fmt.Printf("My 64 bit integer is: %d \n", myInt64)

	// Unsigned Integers
	fmt.Printf("My 8 bit unsigned integer is: %d \n", int8(myUInt8))
	fmt.Printf("My 16 bit Unsigned bit integer is: %d \n", int16(myUInt16))
	fmt.Printf("My 32 bit unsigned bit integer is: %d \n", int32(myUInt32))
	fmt.Printf("My 64 bit unsigned bit integer is: %d \n", int64(myUInt64))
}

func floats() {
	fmt.Println("--Floats--")

	var my32BitFloat float32 = -232
	var my64BitFloat float64 = 232332323232

	fmt.Printf("My 32 bit float is: %v \n", my32BitFloat)
	fmt.Printf("My 64 bit float is: %v \n", my64BitFloat)

}

func bools() {
	fmt.Println("--Booleans--")

	var programRunning bool = true
	var programNotRunning bool = false

	fmt.Printf("Program Running: %v \n", programRunning)
	fmt.Printf("Program Not Running: %v \n", programNotRunning)
}

func strings() {
	fmt.Println("--Strings--")

	var planetName string = "Uranus"

	fmt.Printf("My favorite planet is %v \n", planetName)
}
