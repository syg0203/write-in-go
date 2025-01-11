package main

import(
	"fmt"
	"strconv"
)

func main() {
	var int_var int = 1
	fmt.Printf("%d\n", int_var)

	var str_ascii_var string = string(int_var) // ascii 코드로 변환 ascii 코드 1은 공백
	fmt.Printf("%s\n", str_ascii_var)

	var str_var_1 string = strconv.Itoa(int_var) // 숫자 그대로 문자열 변환
	fmt.Printf("%s\n", str_var_1)

	str_var_2 := fmt.Sprintf("%d", int_var)
	fmt.Printf("%s\n", str_var_2)

	var runeinitialized rune = '💕'
	fmt.Printf("%c\n", runeinitialized)
}

// 변수 선언 방법
// var nameOfVariable variableType = value 
// var nameOfVariable variableType
// nameOfVariable = value
// nameOfVariable := value