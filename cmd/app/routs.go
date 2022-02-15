package main

import "github.com/labstack/echo/v4"

func routs(e *echo.Echo) {
	e.GET("/postjson", getPostJson)
	e.GET("/postxml", getPostXml)
	e.GET("/postsjson", getPostsJson)
	e.GET("/postsxml", getPostsXml)
	e.GET("/comments", getCommByPostjson)
	e.GET("/commentsxml", getCommByPostXml)

	e.POST("/createcomm", saveComm)
	e.POST("/create", createPost)

	e.PUT("/update", updatePost)
	e.PUT("/updatecomm", updateComm)

	e.DELETE("/delete", deletePostById)
	e.DELETE("/deletecomm", deleteComm)
}
