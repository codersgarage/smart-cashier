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

func RegisterDiaryRoutes(g *echo.Group) {
	func(g echo.Group) {
		g.Use(middlewares.AuthUser)
		g.POST("/", createDiary)
		g.GET("/", listDiaries)
		g.PATCH("/:diary_id/", updateDiary)
		g.GET("/:diary_id/", getDiary)
		g.DELETE("/:diary_id/", deleteDiary)
	}(*g)
}

func createDiary(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)

	resp := core.Response{}

	pld, err := validators.ValidateCreateDiary(ctx)

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
	m := &models.Diary{
		ID:        utils.NewUUID(),
		UserID:    userID,
		Name:      pld.Name,
		Type:      pld.Type,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

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

	resp.Status = http.StatusCreated
	resp.Data = m
	return resp.ServerJSON(ctx)
}

func updateDiary(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	diaryID := ctx.Param("diary_id")

	resp := core.Response{}

	pld, err := validators.ValidateUpdateDiary(ctx)

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

	if pld.Name != nil {
		m.Name = *pld.Name
	}
	if pld.Type != nil {
		m.Type = *pld.Type
	}
	m.UpdatedAt = time.Now().UTC()

	if err := dr.UpdateDiary(db, m); err != nil {
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

func getDiary(ctx echo.Context) error {
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

func deleteDiary(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)
	diaryID := ctx.Param("diary_id")

	resp := core.Response{}

	db := app.DB().Begin()

	dr := data.NewDiaryRepository()

	err := dr.DeleteAllEntry(db, diaryID)
	if err != nil {
		db.Rollback()
		log.Log().Errorln(err)

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	err = dr.DeleteAllCategory(db, diaryID)
	if err != nil {
		db.Rollback()
		log.Log().Errorln(err)

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	err = dr.DeleteDiary(db, userID, diaryID)
	if err != nil {
		db.Rollback()
		log.Log().Errorln(err)

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	if err := db.Commit().Error; err != nil {
		log.Log().Errorln(err)

		resp.Title = "Database query failed"
		resp.Status = http.StatusInternalServerError
		resp.Code = errors.DatabaseQueryFailed
		resp.Errors = err
		return resp.ServerJSON(ctx)
	}

	resp.Status = http.StatusNoContent
	return resp.ServerJSON(ctx)
}

func listDiaries(ctx echo.Context) error {
	userID := ctx.Get(utils.UserID).(string)

	resp := core.Response{}

	db := app.DB()

	dr := data.NewDiaryRepository()
	d, err := dr.ListDiaries(db, userID, 0, 100)

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
