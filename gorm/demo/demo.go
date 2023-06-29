package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sanketplus/go-mysql-lock"

)

func main() {
	db, _ := sql.Open("mysql", "root:bigdataNV7U@tcp(10.162.195.168:8270)/gorm")

	locker := gomysqllock.NewMysqlLocker(db)
	for i := 0; i < 6; i++ {
		go funcName(strconv.Itoa(i), locker)
	}
	time.Sleep(20 * time.Second)

}

func funcName(id string, locker *gomysqllock.MysqlLocker) {
	lock, err := locker.ObtainTimeout("foo", 0)
	defer func(lock *gomysqllock.Lock) {
		if lock != nil {
			err := lock.Release()
			if err != nil {
				fmt.Printf("dsadasdsadsad:%s\n", err)
			}
			fmt.Printf("%s: release foo lock finally, %s\n", id, time.Now())
		}
	}(lock)
	if err != nil {
		fmt.Printf("%s: get foo lock failed, %s\n", id, time.Now())
		return
	}
	fmt.Printf("%s: get foo lock, %s\n", id, time.Now())
	r := rand.Int63n(10)
	time.Sleep(time.Duration(r) * time.Second)
	fmt.Printf("%s: release foo lock, %s\n", id, time.Now())
	return
}
