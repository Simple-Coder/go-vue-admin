package main

import "fmt"

func main() {
	////声明
	//var a string = "hello go go "
	//fmt.Println(a)
	//b := "hello go go bbb"
	//fmt.Println(b)
	//fmt.Println(A)
	//fmt.Println("----基本数据类型")
	//
	//var str1 string = "123"
	//atoi, _ := strconv.Atoi(str1)
	//
	//fmt.Println(atoi)
	//fmt.Printf("%T", atoi)

	//数组和切片
	//a := [3]int{0, 2, 1}
	//b := [...]int{0, 2, 1}
	//c := new([10]int)
	//fmt.Println(a, b, c)
	//
	//for i, v := range a {
	//	fmt.Println(i, v)
	//}
	//
	//fmt.Println(len(a), cap(a))

	//切片
	//a := [3]int{0, 1, 3}
	//cl := a[2:]
	//fmt.Println(cl)
	//append
	//cl = append(cl, 5)
	//fmt.Println(cl)
	//copy
	//copy()

	/*
		map
	*/

	//m1 := map[string]string{}
	//m1 := make(map[string]string)
	//m1["name"] = "张三"
	//m1["name1"] = "张三1"
	//fmt.Println(m1, len(m1))
	//
	//delete(m1, "name")
	//fmt.Println(m1, len(m1))
	//
	//for k, v := range m1 {
	//	fmt.Println(k, v)
	//}

	/**
	  func
	*/
	//b := func() {
	//	fmt.Println("我是匿名函数")
	//}
	//b()
	//ar := []string{"123", "falksdjf "}
	//sum(1, ar...)

	//自执行函数
	//(func() {
	//	fmt.Println("我是自治性函数")
	//})()

	//mo()(4)
	defer first()
	fmt.Println(123)
}

func first() {
	fmt.Println("我想最先执行")
}

//闭包函数
func mo() func(int) {
	return func(num int) {
		fmt.Println("闭包函数", num)
	}
}

func sum(a int, b ...string) int {
	for k, v := range b {
		fmt.Println(k, v)
	}
	return 0
}
