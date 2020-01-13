package api

import (
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/core"
	"github.com/codersgarage/smart-cashier/data"
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/codersgarage/smart-cashier/log"
	"github.com/codersgarage/smart-cashier/middlewares"
	"github.com/codersgarage/smart-cashier/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterUserRoutes(g *echo.Group) {
	func(g echo.Group) {
		g.Use(middlewares.AuthUser)
		g.PUT("/", update)
		g.GET("/", get)
	}(*g)
}

func update(ctx echo.Context) error {
	return nil
}

func get(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)

	resp := core.Response{}

	db := app.DB()

	uc := data.NewUserRepository()
	u, err := uc.Get(db, userID)

	if err != nil {
		log.Log().Errorln(err)

		resp.Title = "Failed to get user profile"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Data = map[string]interface{}{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"status":     u.Status,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
	resp.Status = http.StatusOK
	return resp.ServerJSON(ctx)
}
