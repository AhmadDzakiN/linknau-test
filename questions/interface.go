package questions

/*
Interface is a type that has a set of empty methods (method signatures only / without the definition)
Any type that implements all the methods in an interface implicitly satisfies the interface
Usage of interfaces in Go is an example of the abstraction concept in Object-Oriented Programming (OOP) to hide complex implementation of methods
*/

type LivingBeing interface {
	Speak() string
	Run() string
	GetAge() uint
}

type Cat struct {
	Name string
	Age  uint
}

// Speak method implemented by Person struct
func (p Person) Speak() string {
	return "Hello my name is " + p.Name
}

// Run method implemented by Person struct
func (p Person) Run() string {
	return "A person named " + p.Name + " is running"
}

// GetAge method implemented by Person struct
func (p Person) GetAge() uint {
	return p.Age
}

// Speak method implemented by Cat struct
func (c Cat) Speak() string {
	return "Miaw miaw miaw"
}

// Run method implemented by Cat struct
func (c Cat) Run() string {
	return "A cat named " + c.Name + " is running"
}

// GetAge method implemented by Cat struct
func (c Cat) GetAge() uint {
	return c.Age
}
