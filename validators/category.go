package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/codersgarage/smart-cashier/models"
	"github.com/labstack/echo/v4"
)

type ReqCreateCategory struct {
	Name string `json:"name" valid:"required,stringlength(3|100)"`
}

func ValidateCreateCategory(ctx echo.Context) (*ReqCreateCategory, error) {
	pld := &ReqCreateCategory{}

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

type ReqUpdateCategory struct {
	Name string           `json:"name" valid:"required,stringlength(3|100)"`
	Type models.DiaryType `json:"type" valid:"required"`
}

func ValidateUpdateCategory(ctx echo.Context) (*ReqUpdateCategory, error) {
	pld := &ReqUpdateCategory{}

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
