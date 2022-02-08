package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/* Save all posts by userId
and save comments by postId to database from jsonplaceholder */

func saveAllPostsByUserId(c CrudMethods, id int) error {

	chanId := make(chan int, 50)

	var target Post

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", id)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &target); err != nil {
		log.Fatal(err)
	}

	c.Create(target)

	go saveAllCommentsByPostId(chanId)

	for i := 0; i < len(target); i++ {
		chanId <- target[i].ID
	}

	return nil
}

func saveAllCommentsByPostId(id chan int) error {
	db := Connect()

	for chanId := range id {
		url := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments?postId=%d", chanId)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		var target Comments

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &target); err != nil {
			panic(err)
		}

		if err := db.Table("comments").Create(&target).Error; err != nil {
			log.Fatal(err)
			panic(err)
		}

		res.Body.Close()
	}
	return nil
}
