package main

import (
	"net/http"
	"strconv"
)

func AllPosts(w http.ResponseWriter, r *http.Request) {
	var c CrudMethods = &Post{}

	result, err := c.getAll()
	if err != nil {
		panic(err)
	}

	w.WriteHeader(200)
	w.Write(result)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	var c CrudMethods = &Post{}

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	b, err := c.GetById(id)

	if err != nil {
		http.NotFound(w, r)
		// time.Sleep(2 * time.Second)
		// http.Redirect(w, r, "/posts", http.StatusTemporaryRedirect)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

func getCommentsByPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("postid"))
	c := &Comments{}
	b, err := c.getComm(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

func xmlPost(w http.ResponseWriter, r *http.Request) {
	x := XmlPost{}
	b, err := x.getXml()
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

func xmlPostById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var x XmlPost
	b, err := x.getXmlById(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}

func xmlComByPost(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("postid"))
	var x XmlComments

	b, err := x.getXmlComById(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(200)
	w.Write(b)
}
