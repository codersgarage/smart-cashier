package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/codersgarage/smart-cashier/models"
	"github.com/labstack/echo/v4"
)

type ReqCreateDiary struct {
	Name string           `json:"name" valid:"required,stringlength(3|100)"`
	Type models.DiaryType `json:"type" valid:"required"`
}

func ValidateCreateDiary(ctx echo.Context) (*ReqCreateDiary, error) {
	pld := &ReqCreateDiary{}

	if err := ctx.Bind(pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(pld)
	if ok && pld.Type.IsValid() {
		return pld, nil
	}

	ve := errors.ValidationError{}

	if !pld.Type.IsValid() {
		ve.Add("type", "is invalid")
	}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}

type ReqUpdateDiary struct {
	Name *string           `json:"name"`
	Type *models.DiaryType `json:"type"`
}

func ValidateUpdateDiary(ctx echo.Context) (*ReqUpdateDiary, error) {
	pld := &ReqUpdateDiary{}

	if err := ctx.Bind(pld); err != nil {
		return nil, err
	}
	return pld, nil
}
