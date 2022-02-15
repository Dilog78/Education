package handle

import (
	"education/pkg/comment"
	"encoding/xml"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCommByPostIdjson(c echo.Context) error {
	postId, _ := strconv.Atoi(c.QueryParam("postid"))

	comm, err := comment.GetCommByPost(postId)
	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}
	if len(comm) == 0 {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	return c.JSONPretty(http.StatusOK, comm, " ")
}

func GetCommByPostIdXml(c echo.Context) error {
	postId, _ := strconv.Atoi(c.QueryParam("postid"))

	comm, err := comment.GetCommByPost(postId)

	if err != nil {
		http.NotFound(c.Response(), c.Request())
		return err
	}
	if len(comm) == 0 {
		http.NotFound(c.Response(), c.Request())
		return err
	}

	b, _ := xml.MarshalIndent(comm, " ", " ")

	return c.String(200, string(b))
}

func SaveComm(c echo.Context) error {
	postId, _ := strconv.Atoi(c.FormValue("postid"))
	name := c.FormValue("name")
	email := c.FormValue("email")
	body := c.FormValue("body")

	var comm comment.Comment = comment.NewComm(postId, name, email, body)
	id, err := comm.CreateComm()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, id)
}

func UpdateCommById(c echo.Context) error {
	id, _ := strconv.Atoi(c.FormValue("id"))
	name := c.FormValue("name")
	email := c.FormValue("email")
	body := c.FormValue("body")

	var comm comment.Comment = comment.NewUpdateComm(name, email, body)

	err := comm.UpdateComm(id)
	if err != nil {
		return err
	}

	return c.NoContent(204)
}

func DeleteCommById(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	comm := comment.Comment{}
	err := comm.DeleteComm(id)
	if err != nil {
		return err
	}

	return c.NoContent(204)

}
