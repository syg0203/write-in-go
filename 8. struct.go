package main

import "fmt"

type Human struct {
	Name         string
	Email        string
	Income     int
}

func main() {
	var yg Human
	yg.Name = "ygshin"
	yg.Email = "yg@email.com"
	yg.Income = 30000000

	var shin Human = Human{
		"shin",
		"shin@email.com",
		30000001,
	}
	fmt.Println(yg)
	fmt.Println(shin)
}
