package api

import (
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/core"
	"github.com/codersgarage/smart-cashier/data"
	"github.com/codersgarage/smart-cashier/errors"
	"github.com/codersgarage/smart-cashier/log"
	"github.com/codersgarage/smart-cashier/middlewares"
	"github.com/codersgarage/smart-cashier/models"
	"github.com/codersgarage/smart-cashier/utils"
	"github.com/codersgarage/smart-cashier/validators"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func RegisterCategoryRoutes(g *echo.Group) {
	func(g echo.Group) {
		g.Use(middlewares.AuthUser)
		g.POST("/", createCategory)
		g.GET("/", listCategories)
		g.PATCH("/:category_id", updateCategory)
		g.GET("/:category_id", getCategory)
		g.DELETE("/:category_id", deleteCategory)
	}(*g)
}

func createCategory(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	diaryID := ctx.Param("diary_id")

	resp := core.Response{}

	pld, err := validators.ValidateCreateCategory(ctx)

	if err != nil {
		log.Log().Errorln(err)

		resp.Title = "Invalid data"
		resp.Status = http.StatusUnprocessableEntity
		resp.Code = errors.CategoryCreationDataInvalid
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	db := app.DB()

	dr := data.NewDiaryRepository()
	d, err := dr.GetDiary(db, userID, diaryID)
	if err != nil {
		log.Log().Errorln(err)

		if errors.IsRecordNotFoundError(err) {
			resp.Title = "Diary not found"
			resp.Status = http.StatusNotFound
			resp.Code = errors.DiaryNotFound
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	m := &models.Category{
		ID:        utils.NewUUID(),
		DiaryID:   d.ID,
		Name:      pld.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if err := dr.CreateCategory(db, m); err != nil {
		log.Log().Errorln(err)

		msg, ok := errors.IsDuplicateKeyError(err)
		if ok {
			resp.Title = msg
			resp.Status = http.StatusConflict
			resp.Code = errors.DiaryAlreadyExists
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusCreated
	resp.Data = m
	return resp.ServerJSON(ctx)
}

func updateCategory(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	diaryID := ctx.Param("diary_id")

	resp := core.Response{}

	pld, err := validators.ValidateUpdateCategory(ctx)

	if err != nil {
		log.Log().Errorln(err)

		resp.Title = "Invalid data"
		resp.Status = http.StatusUnprocessableEntity
		resp.Code = errors.DiaryCreationDataInvalid
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	db := app.DB()

	dr := data.NewDiaryRepository()
	m, err := dr.GetDiary(db, userID, diaryID)
	if err != nil {
		log.Log().Errorln(err)

		if errors.IsRecordNotFoundError(err) {
			resp.Title = "Diary not found"
			resp.Status = http.StatusNotFound
			resp.Code = errors.DiaryNotFound
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	m.Name = pld.Name
	m.UpdatedAt = time.Now().UTC()

	if err := dr.CreateDiary(db, m); err != nil {
		log.Log().Errorln(err)

		msg, ok := errors.IsDuplicateKeyError(err)
		if ok {
			resp.Title = msg
			resp.Status = http.StatusConflict
			resp.Code = errors.DiaryAlreadyExists
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusOK
	resp.Data = m
	return resp.ServerJSON(ctx)
}

func getCategory(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	diaryID := ctx.Param("diary_id")

	resp := core.Response{}

	db := app.DB()

	dr := data.NewDiaryRepository()
	m, err := dr.GetDiary(db, userID, diaryID)

	if err != nil {
		log.Log().Errorln(err)

		if errors.IsRecordNotFoundError(err) {
			resp.Title = "Diary not found"
			resp.Status = http.StatusNotFound
			resp.Code = errors.DiaryNotFound
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusOK
	resp.Data = m
	return resp.ServerJSON(ctx)
}

func deleteCategory(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	diaryID := ctx.Param("diary_id")

	resp := core.Response{}

	db := app.DB()

	dr := data.NewDiaryRepository()
	err := dr.DeleteDiary(db, userID, diaryID)

	if err != nil {
		log.Log().Errorln(err)

		if errors.IsRecordNotFoundError(err) {
			resp.Title = "Diary not found"
			resp.Status = http.StatusNotFound
			resp.Code = errors.DiaryNotFound
			resp.Errors = err
			return resp.ServerJSON(ctx)
		}

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusNoContent
	return resp.ServerJSON(ctx)
}

func listCategories(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	diaryID := ctx.Param("diary_id")

	resp := core.Response{}

	db := app.DB()

	dr := data.NewDiaryRepository()
	d, err := dr.ListCategories(db, userID, diaryID, 0, 100)

	if err != nil {
		log.Log().Errorln(err)

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusOK
	resp.Data = d
	return resp.ServerJSON(ctx)
}
