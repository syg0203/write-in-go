package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		fmt.Println(s, "***", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// 함수를 동기적으로 실행
	say("Sync1")
	say("Sync2")
	say("Sync3")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)
}
