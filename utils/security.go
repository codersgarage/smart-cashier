package utils

import (
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

const (
	PasswordCost = 7
	UserID       = "user_id"
)

func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost)
	return string(bytes), err
}

func CheckPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func ParseBearerToken(ctx echo.Context) (string, error) {
	bearer := ctx.Request().Header.Get("Authorization")
	bearerWithToken := strings.Split(bearer, " ")

	if len(bearerWithToken) != 2 {
		return "", errors.NewError("Bearer token not found")
	}
	return bearerWithToken[1], nil
}

func GetUserID(ctx echo.Context) string {
	return ctx.Get(UserID).(string)
}
