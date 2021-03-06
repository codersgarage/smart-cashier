package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/labstack/echo/v4"
)

type ReqCreateEntry struct {
	CategoryID string  `json:"category_id" valid:"required,stringlength(1|100)"`
	Note       string  `json:"note" valid:"required,stringlength(1|100)"`
	Amount     float64 `json:"amount" valid:"required"`
}

func ValidateCreateEntry(ctx echo.Context) (*ReqCreateEntry, error) {
	pld := &ReqCreateEntry{}

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

type ReqUpdateEntry struct {
	CategoryID *string `json:"category_id"`
	Note       *string `json:"note"`
	Amount     float64 `json:"amount" valid:"required"`
}

func ValidateUpdateEntry(ctx echo.Context) (*ReqUpdateEntry, error) {
	pld := &ReqUpdateEntry{}

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
