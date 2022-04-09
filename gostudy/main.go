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
	//defer first()
	//fmt.Println(123)

	/*指针*/

	//var a int = 123
	//
	//fmt.Println(a)
	//var b *int
	////取地址
	//b = &a
	////修改值
	//*b = 999
	//fmt.Println(a, b)
	//fmt.Println(a == *b, &a == b)
	//
	//fmt.Println("-------")
	////数组指针和指针数组
	//var arr [5]string
	//arr = [5]string{"1,2,3,4,5"}
	//var arrp *[5]string
	//arrp = &arr
	//
	//fmt.Println(arrp)
	//
	////指针数组
	//var arrpp [5]*string
	//var str1 = "str1"
	//var str2 = "str2"
	//var str3 = "str3"
	//var str4 = "str4"
	//var str5 = "str5"
	//
	//arrpp = [5]*string{&str1, &str2, &str3, &str4, &str5}
	//*arrpp[2] = "33333"
	//fmt.Println(str3)

	//var str1 = "我定义了"
	//fmt.Println(str1)
	//pointFun(&str1)
	//fmt.Println(str1)

	/*结构体*/
	//var user User
	//user.Name = "张三"
	//user.Age = 18
	//user.Sex = true
	//user.hobby = []string{"篮球"}
	//fmt.Println(user)

	//user := User{
	//	Name:  "李四",
	//	Age:   0,
	//	Sex:   true,
	//	hobby: []string{"乒乓球"},
	//}
	//fmt.Println(user)
	//
	//user := new(User)
	//user.Name = "new name"
	//
	//c := user
	//c.Age = 15
	//fmt.Println(user)

	//var user1 *User
	//user1 = &user
	//
	////user1.Name = "修改了"
	//(*user1).Name = "修改了"
	//fmt.Println(user, user1, *user1, &user1)

	//user := User{
	//	Name:  "李四",
	//	Age:   0,
	//	Sex:   true,
	//	hobby: []string{"乒乓球"},
	//	MyHome: Home{
	//		address: "北京",
	//	},
	//}
	//
	//user.Eat("面条")
	//fmt.Println(user)
	//user.MyHome.Open()
}

type User struct {
	Name   string
	Age    int
	Sex    bool
	hobby  []string
	MyHome Home
}

type Home struct {
	address string
}

func (home *Home) Open() {
	fmt.Println("open", home.address)
}
func (user *User) Eat(foodName string) {
	fmt.Printf("%v正在吃%v", user.Name, foodName)
}

func pointFun(p1 *string) {
	*p1 = "我变了"
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
