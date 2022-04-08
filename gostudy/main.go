package main

import (
	"fmt"
	"strconv"

	//"gostudy/testpackage"
	//cool "gostudy/testpackage"
	. "gostudy/testpackage"
)

func main() {
	//声明
	var a string = "hello go go "
	fmt.Println(a)
	b := "hello go go bbb"
	fmt.Println(b)
	fmt.Println(A)
	fmt.Println("----基本数据类型")

	var str1 string = "123"
	atoi, _ := strconv.Atoi(str1)

	fmt.Println(atoi)
	fmt.Printf("%T", atoi)

}
