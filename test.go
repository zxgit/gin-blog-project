/*
@Time : 2021/1/24 18:42
@Author : zhangxin
*/
package main

import (
	"fmt"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
)
type Tag struct{
	Id int
	Name string
}
func main() {
	/*db,err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local","root",	"root",	"127.0.0.1","go-blog"))
	if err!=nil {
		fmt.Println(err)
	}*/
	t:=new(Tag)
	t.Id=222
	t.Name="2222"
	fmt.Println(t)
	//db.Create(t)
}
