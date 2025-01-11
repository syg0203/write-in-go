package main

import "fmt"

func main() {
	// 출력함수 예시 (print)
	// 기본 출력 함수 func Print
	// 공백 포함 출력 함수 func Println
	// 포맷 지정 출력 함수 func Printf
	var x_define int8 = 5
	var y_define int8 = 3
	fmt.Print("x is ", x_define, "\n", "y", "is", y_define, "\n")
	fmt.Println("x is", x_define, "\n", "y", "is", y_define)
	fmt.Printf("x is %d\ny is %d\n",x_define,y_define)

	// 입력함수 예시 (scan)
	// 기본 입력 함수 func Scan
	// 개행 종료 입력 함수 func Scanln
	// 포맷 지정 입력 함수 func Scanf
	var x_input int
	var y_input int
	fmt.Scan(&x_input, &y_input)
	fmt.Printf("x is %d\ny is %d\n",x_input,y_input)
	fmt.Scanln(&x_input, &y_input)
	var z string
	fmt.Scanln(&z)
	fmt.Println(z)
	var hello string
	fmt.Scanf("%s", &hello)
	fmt.Println(hello)
	// 문자열 반환 함수 func Sprint, Sprintf, Sprintln
	// print함수가 출력한 결과를 string타입으로 반환하여 변수에 저장
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n")
	fmt.Println(s)
}