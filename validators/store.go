package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
	"github.com/shopicano/shopicano-backend/errors"
	"github.com/shopicano/shopicano-backend/models"
	"github.com/shopicano/shopicano-backend/utils"
	"time"
)

func ValidateCreateStore(ctx echo.Context) (*models.Store, error) {
	pld := struct {
		Name        string `json:"name" valid:"required,stringlength(1|100)"`
		Address     string `json:"address" valid:"required,stringlength(1|100)"`
		City        string `json:"city" valid:"required,stringlength(1|30)"`
		Country     string `json:"country" valid:"required,stringlength(1|30)"`
		Postcode    string `json:"postcode" valid:"required,stringlength(1|100)"`
		Email       string `json:"email" valid:"required,email"`
		Phone       string `json:"phone" valid:"required,stringlength(1|20)"`
		Description string `json:"description" valid:"required,stringlength(1|1000)"`
		Image       string `json:"image"`
	}{}

	if err := ctx.Bind(&pld); err != nil {
		return nil, err
	}

	ok, err := govalidator.ValidateStruct(&pld)
	if ok {
		return &models.Store{
			ID:                       utils.NewUUID(),
			Name:                     pld.Name,
			Email:                    pld.Email,
			Phone:                    pld.Phone,
			Status:                   models.StoreRegistered,
			Address:                  pld.Address,
			Description:              pld.Description,
			IsOrderCreationEnabled:   false,
			IsProductCreationEnabled: false,
			Postcode:                 pld.Postcode,
			City:                     pld.City,
			Country:                  pld.Country,
			CreatedAt:                time.Now().UTC(),
			UpdatedAt:                time.Now().UTC(),
		}, nil
	}

	ve := errors.ValidationError{}

	for k, v := range govalidator.ErrorsByField(err) {
		ve.Add(k, v)
	}

	return nil, &ve
}