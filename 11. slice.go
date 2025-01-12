package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("slHeader.Data : %v\n", slHeader.Data)
	fmt.Printf("slHeader.Len : %v\n", slHeader.Len)
	fmt.Printf("slHeader.Cap : %v\n", slHeader.Cap)
	fmt.Println(slice)
	slice = append(slice, 6, 7, 8)
	slHeader = (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	fmt.Printf("slHeader.Data : %v\n", slHeader.Data)
	fmt.Printf("slHeader.Len : %v\n", slHeader.Len)
	fmt.Printf("slHeader.Cap : %v\n", slHeader.Cap)
	fmt.Println(slice)
}
