package sts

import "fmt"

type Person struct {
	name    string
	age     int8
	address Address
}

type Address struct {
	title string
}

type Animal struct {
	name  string
	color string
}

func Structs() {
	fmt.Println("Structs")

	var person1 Person

	person1.name = "Johnson"
	person1.name = "John"
	person1.age = 111
	person1.address = Address{"this is my address ya'll"}

	var animal1 Animal = Animal{name: "frog", color: "green"}
	animal2 := Animal{"goat", "brown"}

	fmt.Println(animal1)
	fmt.Println(animal2)
	passTo(person1)

	person1.printName()

	person1.changeName("rufus")

	person1.printName()

	object := struct{ id string }{id: "ifadlfjalfa flafjalfjdfkjdalfjalfja"}

	println(object.id)
	println("heyo")
}

func passTo(person Person) {
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.address)
}

func (person Person) printName() {
	fmt.Println("Person Name:")
	fmt.Println(person.name)
}

func (person *Person) changeName(name string) {
	person.name = name
}
