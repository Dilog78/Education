package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {

	var crud CrudMethods = &Post{}
	saveAllPostsByUserId(crud, 1)
	crud.DeleteById([]int{1})
	crud.UpdateById(2, "body", "HI")

	http.HandleFunc("/posts", AllPosts)
	http.HandleFunc("/posts/post", GetPostById)     // ?id=...
	http.HandleFunc("/comments", getCommentsByPost) // ?postid=...
	http.HandleFunc("/xmlposts", xmlPost)
	http.HandleFunc("/xmlposts/post", xmlPostById) // ?id=...
	http.HandleFunc("/xmlcomments", xmlComByPost)  // ?postid=...

	fmt.Println("Number of runnable goroutines: ", runtime.NumGoroutine())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
