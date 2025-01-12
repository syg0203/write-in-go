// 포인터 사용 이유 : 퍼포먼스를 고려할때 인스턴스(메모리의 실체)을 가져오는게 아닌 포인터(메모리 주소)만 가져온다.

package main

import "fmt"

func main(){
	var int_val int = 8
	var p_val *int = &int_val // &는 인스턴스의 메모리 주소값
	fmt.Println(p_val)
	fmt.Println(*p_val) // *는 주소가 가르키는 인스턴스 값을 나타냄 
}