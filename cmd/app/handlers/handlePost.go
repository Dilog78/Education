package handle

import (
	middleware "education/pkg/jwtMiddleware"
	"education/pkg/post"
	"encoding/xml"
	"log"
	"net/http"
	"strconv"

	_ "education/docs"

	"github.com/labstack/echo/v4"
	_ "github.com/swaggo/echo-swagger"
)

// CreateTodo godoc
// @Summary      return post by id
// @Tags root
// @Produce      json
// @Param        id query string true "get postbyId"
// @Success 204 {null} nil "Your message here"
// @Failure      404   {string}   string  "ok"
// @Router       /postjson [get]

func GetPostJson(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	p := post.Post{}
	_, err := p.GetPost(id)
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	return c.JSONPretty(http.StatusOK, p, " ")
}

func GetPostXml(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	p := post.Post{}
	_, err := p.GetPost(id)
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	return c.XMLPretty(http.StatusOK, p, " ")
}

func GetPostsJson(c echo.Context) error {

	id, err := middleware.ParseAuth(c)
	if err != nil {
		log.Fatal(err)
	}

	p, err := post.GetPosts(id)
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	return c.JSONPretty(http.StatusOK, p, " ")
}

func GetPostsXml(c echo.Context) error {
	id, err := middleware.ParseAuth(c)
	if err != nil {
		log.Fatal(err)
	}
	p, err := post.GetPosts(id)
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	b, _ := xml.MarshalIndent(p, " ", " ")
	return c.String(http.StatusOK, string(b))
}

func CreatePost(c echo.Context) error {

	userId, _ := strconv.Atoi(c.FormValue("userid"))
	title := c.FormValue("title")
	body := c.FormValue("body")

	var p post.Post = *post.NewPost(userId, title, body)

	id, err := p.Create()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, id)
}

func DeletePostById(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))

	p := post.Post{}
	err := p.DeletePost(id)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

func UpdatePostById(c echo.Context) error {
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
