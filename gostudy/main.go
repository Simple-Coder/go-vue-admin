package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

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

	/**
	接口  Interface
	*/
	//c := Cat{
	//	Name: "Tom",
	//	Sex:  false,
	//}
	//MyFunc(c)

	/**
	goroutine channel通讯桥梁
	*/

	//var wg sync.WaitGroup
	//wg.Add(1)
	//go Run(&wg)
	//wg.Wait()

	//无缓冲区channel
	//c1 := make(chan int)
	//有缓冲区channel
	//c1 := make(chan int, 5)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		c1 <- i
	//	}
	//}()
	//
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-c1)
	//}

	//c1 := make(chan int, 5)
	//var readc <-chan int = c1
	//var writec chan<- int = c1
	//writec <- 1
	//fmt.Println(<-readc)

	//c1 := make(chan int, 5)
	//c1 <- 1
	//c1 <- 2
	//c1 <- 3
	//c1 <- 4
	//c1 <- 5
	//close(c1) //用完关闭
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)
	//fmt.Println(<-c1)

	//c1 := make(chan int, 5)
	//c1 <- 1
	//c1 <- 2
	//c1 <- 3
	//c1 <- 4
	//c1 <- 5
	//close(c1) //先close 才能用for range
	//for v := range c1 {
	//	fmt.Println(v)
	//}

	//c1 := make(chan int, 5)
	//c2 := make(chan int, 5)
	//c3 := make(chan int, 5)
	//
	//c1 <- 123
	//c2 <- 123
	//c3 <- 123
	////随机执行多个
	//select {
	//case <-c1:
	//	fmt.Println("c1")
	//case <-c2:
	//	fmt.Println("c2")
	//case <-c3:
	//	fmt.Println("c3")
	//default:
	//	fmt.Println("都不满足")
	//}

	//c := make(chan int)
	//var readc <-chan int = c
	//var writec chan<- int = c
	//
	//go SetChan(writec)
	//
	//GetChan(readc)

	/**
	断言和反射
	*/
	//stu := Student{
	//	Name: "张三",
	//	Age:  0,
	//	School: School{
	//		"鹿泉一中",
	//	},
	//}
	//check(&stu)
	SyncClass()
}
func SyncClass() {
	//1.互斥锁 mutex
	//mutex := &sync.Mutex{}
	//go lockFun(mutex)
	//go lockFun(mutex)
	//go lockFun(mutex)
	//go lockFun(mutex)
	//for {
	//
	//}

	//2.读写
	//mutex := &sync.RWMutex{}
	//go lockFun(mutex)
	//go lockFun(mutex)
	//go lockReadFun(mutex)
	//go lockReadFun(mutex)
	//go lockReadFun(mutex)
	//go lockReadFun(mutex)
	//for {
	//}

	////3. once 只会执行一次
	//once := &sync.Once{}
	//for i := 0; i < 5; i++ {
	//	once.Do(func() {
	//		fmt.Println("我执行了")
	//	})
	//}

	//4. waitGroup
	//group := &sync.WaitGroup{}
	//group.Add(2)
	//
	//go func() {
	//	time.Sleep(3 * time.Second)
	//	fmt.Println("hp -1")
	//	group.Done()
	//}()
	//go func() {
	//	time.Sleep(6 * time.Second)
	//	fmt.Println("hp -1")
	//	group.Done()
	//}()
	//
	//group.Wait()
	//fmt.Println("hp==0 over")

	//5.map
	//m := &sync.Map{}
	//go func() {
	//	for {
	//		m.Store(1, 1)
	//	}
	//}()
	//go func() {
	//	for {
	//		fmt.Println(m.Load(1))
	//	}
	//}()
	//time.Sleep(100)

	//m.Store(1, 1)
	//m.Store(2, 2)
	//m.Store(3, 3)
	//m.Range(func(key, value interface{}) bool {
	//	fmt.Println(key, value)
	//	time.Sleep(1 * time.Second)
	//	return true
	//})

	//并发池
	//pool := &sync.Pool{}
	//pool.Put(1)
	//pool.Put(2)
	//pool.Put(3)
	//pool.Put(4)
	//for i := 0; i < 4; i++ {
	//	fmt.Println(pool.Get())
	//}

	//等待唤醒机制
	//co := sync.NewCond(&sync.Mutex{})
	//go func() {
	//	co.L.Lock()
	//	fmt.Println("lock1")
	//	co.Wait()
	//	fmt.Println("unlock1")
	//	co.L.Unlock()
	//}()
	//go func() {
	//	co.L.Lock()
	//	fmt.Println("lock2")
	//	co.Wait()
	//	fmt.Println("unlock2")
	//	co.L.Unlock()
	//}()
	//
	//time.Sleep(2 * time.Second)
	////co.Broadcast()
	//co.Signal()
	//co.Signal()
	//time.Sleep(2 * time.Second)

	/**
	文件操作
	*/

	//file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//for true {
	//	defer file.Close()
	//	bytes := make([]byte, 12)
	//	read, err := file.Read(bytes)
	//	if err != nil {
	//		return
	//	}
	//	fmt.Println(string(bytes), read)
	//}
	//追加
	//file.Seek(0, io.SeekEnd)
	//file.Write([]byte("1234566788"))

	//
	//defer file.Close()
	//reader := bufio.NewReader(file)
	//reader.ReadLine()

	//defer file.Close()
	//all, err := ioutil.ReadAll(file)
	//fmt.Println(string(all))
	//dir, err := os.ReadDir("./")
	//fmt.Println(dir, err)
	//for _, entry := range dir {
	//	fmt.Println(entry.Name())
	//	fmt.Println(entry.Type())
	//}

	//缓冲区
	//writer := bufio.NewWriter(file)
	//writer.WriteString()
	//writer.Flush()
	//reader := bufio.NewReader(file)

	//复制文件

	//从头覆盖操作
	//io.Copy()

}
func lockReadFun(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("疯狂治疗。。")
	time.Sleep(1 * time.Second)
	lock.RUnlock()
}

//断言和反射 开始
func check(obj interface{}) {
	//v.(Student).SaySchool("鹿泉一中")
	//switch v.(type) {
	//case User:
	//	fmt.Printf("user")
	//case Student:
	//	fmt.Printf("student")
	//
	//}

	//t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	//for i := 0; i < t.NumField(); i++ {
	//	fmt.Println(v.Field(i))
	//}
	//
	//fmt.Println(t, v)
	//fmt.Println(v.FieldByIndex([]int{0}))
	//fmt.Println(v.FieldByIndex([]int{2, 0}))
	//fmt.Println(v.FieldByName("Name"))

	//ty := t.Kind()
	//if ty == reflect.Struct {
	//	fmt.Println("我是struct")
	//}
	//if ty == reflect.String {
	//
	//}

	elem := v.Elem()
	elem.FieldByName("Name").SetString("张三三=")
	fmt.Println(obj)
}

type Student struct {
	Name string
	Age  int
	School
}
type School struct {
	Name string
}

func (stu Student) SaySchool(name string) {
	fmt.Println("我的学校是", name)
}

//断言和反射 结束

//goroutine channel开始

func Run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了")
	wg.Done()
}
func GetChan(readc <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("get函数取到set函数返回信息：%d \n", <-readc)
	}
}
func SetChan(writec chan<- int) {
	for i := 0; i < 10; i++ {
		writec <- i
	}
}

//goroutine channel结束

//----接口开始

type Animal interface {
	Eat()
	Run()
}
type Cat struct {
	Name string
	Sex  bool
}
type Dog struct {
	Name string
	Sex  bool
}

func (cat Cat) Eat() {
	fmt.Println(cat.Name, "开始吃")
}
func (cat Cat) Run() {
	fmt.Println(cat.Name, "开始跑")
}
func (dog Dog) Eat() {
	fmt.Println(dog.Name, "开始吃")
}
func (dog Dog) Run() {
	fmt.Println(dog.Name, "开始跑")
}

func MyFunc(animal Animal) {
	animal.Run()
	animal.Eat()
}

//----接口结束

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
