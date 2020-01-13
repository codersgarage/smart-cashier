package middlewares

import (
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/core"
	"github.com/codersgarage/smart-cashier/data"
	"github.com/codersgarage/smart-cashier/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

var AuthUser = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		resp := core.Response{}

		token, err := utils.ParseBearerToken(ctx)
		if err != nil {
			resp.Status = http.StatusUnauthorized
			resp.Title = "Unauthorized request"
			return resp.ServerJSON(ctx)
		}

		db := app.DB()

		uc := data.NewUserRepository()
		session, err := uc.GetSession(db, token)
		if err != nil {
			resp.Status = http.StatusUnauthorized
			resp.Title = "Unauthorized request"
			return resp.ServerJSON(ctx)
		}

		ctx.Set(utils.UserID, session.UserID)
		return next(ctx)
	}
}
