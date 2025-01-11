package main

import "fmt"

func if_example(input int64) int64 {
	if input < 0 {
		return -input
	}
	return input
}

func switch_example(name string) int{
	switch name {
	case "shin":
		return 21
	case "lee":
		return 22
	default:
		return 0
	}
}

func switch_example2(name string) int{
	switch {
	case name == "shin":
		return 21
	case name == "lee":
		return 22
	default:
		return 0
	}
}

func switch_checktype() {
	var name interface{} = "shin"
	switch t := name.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}


func main(){
	fmt.Println(if_example(-1))
	fmt.Println(switch_example("shin"))
	fmt.Println(switch_example("lee"))
	fmt.Println(switch_example("park"))
	fmt.Println(switch_example2("shin"))
	fmt.Println(switch_example2("lee"))
	fmt.Println(switch_example2("park"))
	switch_checktype()
}