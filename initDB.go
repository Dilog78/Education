package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@/task"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
