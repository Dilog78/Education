package main

import (
	handle "education/cmd/app/handlers"

	"github.com/labstack/echo/v4"
)

func routs(e *echo.Echo) {
	e.GET("/postjson", handle.GetPostJson)
	e.GET("/postxml", handle.GetPostXml)
	e.GET("/postsjson", handle.GetPostsJson)
	e.GET("/postsxml", handle.GetPostsXml)
	e.GET("/comments", handle.GetCommByPostIdjson)
	e.GET("/commentsxml", handle.GetCommByPostIdXml)

	e.POST("/createcomm", handle.SaveComm)
	e.POST("/create", handle.CreatePost)
	e.POST("/createuser", handle.CreateUser)
	e.POST("/session", handle.SessionsUser)

	e.PUT("/update", handle.UpdatePostById)
	e.PUT("/updatecomm", handle.UpdateCommById)

	e.DELETE("/delete", handle.DeletePostById)
	e.DELETE("/deletecomm", handle.DeleteCommById)
}
