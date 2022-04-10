package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"primary_key;column:user_name;type:varchar(100)"`
	Age  int
}

//自定义表名
func (user User) TableName() string {
	return "go_user"
}
func main() {
	db, err := gorm.Open("mysql", "root:123456@(www.xdblog.site:3307)/go_demos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	/// preload 预加载
}
