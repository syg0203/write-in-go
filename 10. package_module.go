package main

import (
	"fmt"
	"local_package/common" // go mod init local_package (모듈명 설정)
)

func main() {
	fmt.Println(common.Sum_yg(1, 2))
	common.Print_name()
}
