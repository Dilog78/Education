package pkg

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:root@/task"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
