package auth

import (
	"education/pkg"
	"time"

	"github.com/dgrijalva/jwt-go"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

const mySigningKey = "Al#%lYrsdgerou453342rB&^%ase"

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Hash     string `json:"-"`
}

func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
		IssuedAt:  time.Now().Unix(),
	})

	return token.SignedString([]byte(mySigningKey))
}

func (u *User) Validator() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Length(6, 100), validation.NilOrNotEmpty),
	)
}

func (u *User) HashPass() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func (u *User) Save() error {
	db := pkg.InitDB()

	if err := db.Table("users").Create(u).Error; err != nil {
		return err
	}
	return nil
}
