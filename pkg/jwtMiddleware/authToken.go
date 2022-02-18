package middleware

import (
	auth "education/pkg/user"
	"errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func ParseAuth(c echo.Context) (string, error) {
	header := c.Request().Header.Get("Authorization")

	if header == "" {
		return "", errors.New("empty header")
	}

	headerParts := strings.Split(header, " ")

	return tokenParse(headerParts[1])
}

func tokenParse(t string) (string, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return auth.MySigningKey, nil
	})
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)

	return claims["jti"].(string), nil
}
