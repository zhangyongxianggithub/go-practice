package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	ID           uint      `gorm:"column:id;type:unit;primaryKey" json:"id"`
	ISBN         string    `gorm:"column:isbn;type:string" json:"isbn"`
	Title        string    `gorm:"column:title" json:"title"`
	Author       string    `gorm:"column:author" json:"author"`
	Description  string    `gorm:"column:description" json:"description"`
	CreatedTime  time.Time `gorm:"column:created_time" json:"createdTime"`
	ModifiedTime time.Time `gorm:"column:modified_time" json:"modifiedTime"`
}

func (*Book) TableName() string {
	return "book"
}

func main() {

	dsn := "root:163766@tcp(a.bestzyx.com:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Read
	var book Book
	db.First(&book, 4) // 根据整型主键查找
	bookValue, _ := json.Marshal(book)
	fmt.Printf("%s", bookValue)
	db.First(&book, "isbn = ?", "9787805019765") // 查找 code 字段值为 D42 的记录
	bookValue, _ = json.Marshal(book)
	fmt.Printf("%s", bookValue)

}
