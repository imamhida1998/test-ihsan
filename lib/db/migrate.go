package db

import (
	"fmt"
	"test-ihsan/model"
)

func Migrate() {
	if exist := Mysql.HasTable("nasabah"); !exist {
		fmt.Println("migrate table nasabah")
		err := Mysql.CreateTable(&model.Nasabah{})
		if err == nil {
			fmt.Println("success migrate table nasabah")
		} else {
			fmt.Println("fail migrate table nasabah: ", err.Error)
		}
	}
}
