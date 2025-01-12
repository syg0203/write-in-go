// private 함수의 경우 대문자로 시작
package main

import "fmt"

func nameOfFunc(a, b int) bool {
	isTrue := (a == b)
	return isTrue
}

func main() {
	isOk := nameOfFunc(3, 5)
	fmt.Println(isOk)
}
// func : function(함수)의 약자로, 함수를 정의할 때 꼭 포함해야 함
// NameOfFunction : 함수의 이름이며, 영대문자로 시작하기 때문에 외부 패키지에서도 사용이 가능함
// (a int, b int) : 함수 내에서 사용될 매개변수 a와 b를 각각 int로 선언함
// (isTrue bool) : 함수의 반환값을 변수 isTrue로 하며, 해당 변수는 bool 타입임
// { ... } : 함수의 내용을 적는 파트이며, 자유롭게 작성함
// isTrue = (a == b) : 비교 연산자 '=='으로 변수 a와 b가 동일한지 체크함
// a == b 가 참일 때는 true, 거짓일 때는 false를 반환함
// return : 값을 반환함 -> 반환할 값을 변수 isTrue로 지정했기 때문에 변수명을 적지 않아도 됨