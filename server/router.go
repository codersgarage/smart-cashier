package server

import (
	"github.com/codersgarage/smart-cashier/api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var router = echo.New()

// getRouter returns the api router
func getRouter() http.Handler {
	router.Use(middleware.Logger())
	//router.Use(middleware.Recover())
	router.Pre(middleware.AddTrailingSlash())
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	router.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})

	registerV1Routes()

	return router
}

func registerV1Routes() {
	v1 := router.Group("/v1")

	userGroup := v1.Group("/users")
	diaryGroup := v1.Group("/diaries")
	diaryEntryCategoriesGroup := diaryGroup.Group("/diaries/:diary_id/categories")
	diaryEntriesGroup := diaryGroup.Group("/diaries/:diary_id/entries")

	fsGroup := v1.Group("/fs")

	api.RegisterLegacyRoutes(v1)
	api.RegisterUserRoutes(userGroup)
	api.RegisterFSRoutes(fsGroup)
	api.RegisterDiaryRoutes(diaryGroup)
	api.RegisterDiaryEntryRoutes(diaryEntriesGroup)
	api.RegisterDiaryEntryCategoryRoutes(diaryEntryCategoriesGroup)
}
