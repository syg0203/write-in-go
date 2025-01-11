package main

import "fmt"

func main(){
	for i := 0; i<100; i++ {
		fmt.Printf("%d ", i)
	}
	println("")
	i := 100
	for i < 200 {
		fmt.Printf("%d ", i)
		i++
	}
	println("")
	var array [4] string
	array[0] = "s"
	array[1] = "h"
	array[2] = "i"
	array[3] = "n"
	for i, v := range array {
		fmt.Printf("index : %d \t value : %s \n", i, v)
	}
}