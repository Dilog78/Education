package handle

import (
	"education/pkg"
	"education/pkg/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	email := c.FormValue("email")
	pass := c.FormValue("password")

	u := auth.User{
		Email:    email,
		Password: pass,
	}

	if err := u.Validator(); err != nil {
		return echo.ErrBadRequest

	}

	h, err := u.HashPass()
	if err != nil {
		return err
	}

	u.Hash = h

	if err := u.Save(); err != nil {
		return err
	}

	u.Password = ""

	return c.JSONPretty(http.StatusCreated, u, " ")
}

func SessionsUser(c echo.Context) error {
	db := pkg.InitDB()
	email := c.FormValue("email")
	pass := c.FormValue("password")

	u := auth.User{
		Email:    email,
		Password: pass,
	}
	if err := db.Table("users").Where("email", u.Email).First(&u).Error; err != nil {
		return echo.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(pass)); err != nil {
		return echo.ErrUnauthorized
	}

	u.Password = ""
	return c.JSONPretty(http.StatusOK, u, " ")
}
