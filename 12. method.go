package main

import "fmt"

type Parameter struct {
	XX string
}

func nameOfFunction(p *Parameter) {
	p.XX = "YY"
}

type Receiver struct {
	XX string
}

func (r *Receiver) nameOfMethod() {
	r.XX = "YY"
}

// 메서드는 함수와 큰 차이가 없으나, 객체에 속한다는 특징
func main() {
	r := Receiver{}
	r.XX = "XX"
	r.nameOfMethod() // 매서드 실행
	fmt.Println(r)

	p := Parameter{}
	p.XX = "XX"
	nameOfFunction(&p) // 함수 실행
	fmt.Println(p)
}
