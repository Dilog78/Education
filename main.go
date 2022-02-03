package main

import (
	"fmt"
	"runtime"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:root@/task"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	postId := make(chan int, 10)
	go savePosts(db, postId)
	go saveComments(db, postId)

	fmt.Println("Number of runnable goroutines: ", runtime.NumGoroutine())
	var input string
	fmt.Scanln(&input)
}
