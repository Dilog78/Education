package handle

import (
	"education/pkg"
	"education/pkg/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c echo.Context) error {
	db := pkg.InitDB()
	u := auth.User{}

	c.Bind(&u)

	if err := db.Table("users").Where("email", u.Email).First(&u).Error; err == nil {
		return c.String(http.StatusUnauthorized, "Email exist")
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

	tokenString, err := u.GenerateToken()
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}

func SessionsUser(c echo.Context) error {
	db := pkg.InitDB()

	u := auth.User{}

	c.Bind(&u)

	pass := u.Password

	if err := db.Table("users").Where("email", u.Email).First(&u).Error; err != nil {
		return echo.ErrUnauthorized
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Hash), []byte(pass)); err != nil {
		return echo.ErrUnauthorized
	}

	tokenString, err := u.GenerateToken()
	if err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": tokenString,
	})
}
