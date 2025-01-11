package main // hello.go 파일은 main 패키지임을 알립니다.
// main 패키지는 진입점이며, main 패키지 파일 내에는 리턴이 없는 main 함수가 필요합니다.

import "fmt" // golang standard library 중 fmt 표준 라이브러리를 포함합니다.

func main() { // main 함수를 작성합니다. 패키지를 초기화한 후에 가장 먼저 실행되는 함수입니다.
	fmt.Println("Hello World") // fmt는 formatted I/O으로, C언어의 stdio와 같은 표준 입출력 라이브러리입니다. Println은 각각의 피연산자 사이에 공백을 추가하고, 가장 마지막에 줄바꿈을 포함합니다. fmt.Print는 공백을 추가하지 않으며, 줄바꿈도 하지 않습니다.
}

// 코드 실행 방법
// go run hello.go : main 패키지 파일을 컴파일 후에 바이너리 파일을 남기지 않고 메모리 위에서 바로 실행합니다.
// go build hello.go : main 패키지 파일을 import 시킨 모든 라이브러리, 연결된 패키지와 함께 컴파일하며 바이너리 파일을 생성합니다.
// ./hello : 빌드된 hello 바이너리 파일을 실행합니다.