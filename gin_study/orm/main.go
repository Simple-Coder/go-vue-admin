package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
	Sex  bool
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(www.xdblog.site:3307)/go_demos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	//create
	//db.Create(&User{
	//	Model: gorm.Model{},
	//	Name:  "李四",
	//	Age:   20,
	//	Sex:   true,
	//})

	//query

	//var u1 User
	//db.First(&u1, "name = ?", "张三")

	var us []User
	//db.Find(&us, "name = ?", "张三")
	//db.Where("name = ?", "张三").Find(&us)
	//db.Find(&us)

	//update

	//db.Where("id = ?", 1).Find(&User{}).Update("name", "王木头")
	//db.Where("id = ?", 1).Find(&User{}).Updates(User{
	//	Name: "张三111",
	//	Age:  100,
	//})

	//delete
	//db.Delete(&User{}, "id = ?", 1)
	//硬删除
	//db.Where("id = ?", 1).Unscoped().Delete(&User{})
	fmt.Println(us)

}
