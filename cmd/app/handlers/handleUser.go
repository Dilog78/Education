package handle

import (
	"education/pkg"
	authgoogle "education/pkg/OAuth2"
	auth "education/pkg/user"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	html := `<html><body>
		<form action="http://localhost:8080/signin" method="POST">
			<p><input name="email"> <input name="password"></p>
			<p><input type="submit"></p>
		   </form>
		   <a href="http://localhost:8080/googleauth"><input type="button" value="гугл авторизация"></a>
	</body></html>`
	fmt.Fprintf(c.Response().Writer, "%s", html)
	return nil
}

func GoogleAuth(c echo.Context) error {
	u := authgoogle.Init()
	return c.Redirect(http.StatusTemporaryRedirect, u)
}

func GoogleCallback(c echo.Context) error {
	content, err := authgoogle.GetUserInfo(c.FormValue("state"), c.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	return c.JSON(200, string(content))
}

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
	u.Password = ""
	return c.JSON(http.StatusCreated, u)
}

func SignIn(c echo.Context) error {
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
	u.Password = ""

	return c.JSON(http.StatusOK, echo.Map{"token": tokenString})
}
