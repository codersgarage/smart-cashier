package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/labstack/echo/v4"
)

func ValidateLogin(ctx echo.Context) (string, string, error) {
	ul := struct {
		Email    string `json:"email" valid:"required,stringlength(3|100)"`
		Password string `json:"password"`
	}{}

	if err := ctx.Bind(&ul); err != nil {
		return "", "", err
	}

	ok, err := govalidator.ValidateStruct(&ul)
	if ok {
		return ul.Email, ul.Password, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return "", "", &ve
}

type ReqRegister struct {
	Name           string  `json:"name" valid:"required,stringlength(3|100)"`
	Email          string  `json:"email" valid:"required,email"`
	ProfilePicture *string `json:"profile_picture"`
	Password       string  `json:"password" valid:"required,stringlength(8|100)"`
}

func ValidateRegister(ctx echo.Context) (*ReqRegister, error) {
	pld := ReqRegister{}

	if err := ctx.Bind(&pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(&pld)
	if ok {
		return &pld, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}
