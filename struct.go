package main

type XmlComments struct {
	PostID int    `xml:"postId"`
	ID     int    `xml:"id"`
	Name   string `xml:"name"`
	Email  string `xml:"email"`
	Body   string `xml:"body"`
}

type XmlPost struct {
	UserID int    `xml:"userId"`
	ID     int    `xml:"id"`
	Title  string `xml:"title"`
	Body   string `xml:"body"`
}

type Post []struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comments []struct {
	PostID int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
