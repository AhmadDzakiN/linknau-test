package main

import (
	"fmt"
	"recruitment-test/linknau-test/questions"
)

func main() {
	// Number 1, You can check questions/struct.go file for more details
	p := questions.Person{
		Name: "Ahmad",
		Age:  26,
	}

	// Print p variable
	fmt.Println(p)

	// Number 2, You can check questions/interface.go file for more details
	c := questions.Cat{
		Name: "Oyen",
		Age:  2,
	}

	fmt.Println("------SPEAK------")
	fmt.Println(p.Speak())
	fmt.Println(c.Speak())

	fmt.Println("------RUN------")
	fmt.Println(p.Run())
	fmt.Println(c.Run())

	fmt.Println("------GETAGE------")
	fmt.Println(p.GetAge())
	fmt.Println(c.GetAge())

	// Number 3, You can check the answers at questions/package_management.go

	// Number 4, You can check questions/auth file for more details
	/*
		 Below command will run a server at localhost:8080
		 There are 2 endpoints available to use ("/" and "/create")
		 Endpoint "/" is a GET endpoint with Bearer Token authorization header and will return a message and the JWT itself
		 Endpoint "/create"/ is a POST endpoint without any authorization. However, it needs a request body "role" to
		create role authorization based on that request body input. The role will be used in endpoint "/" that use role authorization
	*/
	questions.RunNewServer()

	// Number 5, You can check number_five.go file for the function/method implementation. Then you can check number_five_test.go file for the unit test implementation
	// You can run go test -v ./... command in your terminal to run the unit tests
}
