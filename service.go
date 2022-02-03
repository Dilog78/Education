package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
)

func savePosts(db *gorm.DB, postId chan int) {

	type Post []struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?userId=7")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var target Post

	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &target); err != nil {
		log.Fatal(err)
		panic(err)
	}

	if err := db.Table("posts").Create(&target).Error; err != nil {
		log.Fatal(err)
		panic(err)
	}

	for i := 0; i < len(target); i++ {
		postId <- target[i].ID
	}
	close(postId)
}

func saveComments(db *gorm.DB, postId chan int) {

	type Comments []struct {
		PostId int    `json:"postId"`
		Id     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}

	for chanId := range postId {

		urlId := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", chanId)
		resp, err := http.Get(urlId)
		if err != nil {
			log.Fatal(err)
		}

		var target Comments
		body, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &target); err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()

		if err := db.Table("comments").Create(&target).Error; err != nil {
			log.Fatal(err)
			panic(err)
		}
	}
}
