package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
与另一个模型建立一对多的连接，可以有0+个关联模型
*/
// User 有多张 CreditCard，UserID 是外键
type User struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer;"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
}

func main() {

	dsn := "root:163766@tcp(a.bestzyx.com:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &CreditCard{})
	// db.Create(&User{Model: gorm.Model{
	// 	ID: 1,
	// }})
	users, _ := GetAll(db)
	str, _ := json.Marshal(users)
	fmt.Printf("%s", str)

}

// 检索用户列表并预加载信用卡
func GetAll(db *gorm.DB) ([]User, error) {
	var users []User
	err := db.Model(&User{}).Preload("CreditCards").Find(&users).Error
	return users, err
}
