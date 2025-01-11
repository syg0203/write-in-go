package main

import "fmt"

func main(){
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := [3]bool{true, false, true}
	fmt.Println(b)
	var c = [5]float64{1: 1.1, 3: 3.3} // 특정 인덱스값 초기화
	fmt.Println(c)
	d := [...]string{"one", "two", "three"}
	fmt.Println(len(d))
	var mul [2][3]int = [2][3]int{
		{1, 2, 3}, 
		{4, 5, 6},
	}
	fmt.Println(mul)
}