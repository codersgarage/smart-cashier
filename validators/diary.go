package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/labstack/echo/v4"
)

type ReqCreateDiary struct {
	Name string `json:"name" valid:"required,stringlength(3|100)"`
}

func ValidateCreateDiary(ctx echo.Context) (*ReqCreateDiary, error) {
	pld := &ReqCreateDiary{}

	if err := ctx.Bind(pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(pld)
	if ok {
		return pld, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}

type ReqUpdateDiary struct {
	Name string `json:"name" valid:"required,stringlength(3|100)"`
}

func ValidateUpdateDiary(ctx echo.Context) (*ReqUpdateDiary, error) {
	pld := &ReqUpdateDiary{}

	if err := ctx.Bind(pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(&pld)
	if ok {
		return pld, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}
