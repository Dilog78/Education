package main

import (
	"education/pkg/post"
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getPostJson(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	p := post.Post{}
	_, err := p.GetPost(id)
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	return c.JSONPretty(http.StatusOK, p, " ")
}

func getPostXml(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	p := post.Post{}
	_, err := p.GetPost(id)
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	return c.XMLPretty(http.StatusOK, p, " ")
}

func getPostsJson(c echo.Context) error {
	p, err := post.GetPosts()
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	return c.JSONPretty(http.StatusOK, p, " ")
}

func getPostsXml(c echo.Context) error {
	p, err := post.GetPosts()
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	b, _ := xml.MarshalIndent(p, " ", " ")
	return c.String(http.StatusOK, string(b))
}

func createPost(c echo.Context) error {

	userId, _ := strconv.Atoi(c.FormValue("userid"))
	title := c.FormValue("title")
	body := c.FormValue("body")

	var p post.Post = *post.NewPost(userId, title, body)

	id, err := p.CreatePost()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, id)
}

func deletePostById(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	p := post.Post{}
	err := p.DeletePost(id)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

func updatePost(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("ID"))
	title := c.FormValue("Title")
	body := c.FormValue("Body")

	var p post.Post = *post.NewUpdate(title, body)

	err := p.UpdatePost(id)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}
