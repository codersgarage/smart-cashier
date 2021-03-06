package tasks

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/data"
	"github.com/codersgarage/smart-cashier/log"
	"github.com/codersgarage/smart-cashier/services"
	"github.com/codersgarage/smart-cashier/utils"
	"time"
)

const (
	SendSignUpVerificationEmailTaskName = "send_sign_up_verification_email"
)

func SendSignUpVerificationEmailFn(userID string) error {
	db := app.DB().Begin()

	userDao := data.NewUserRepository()
	u, err := userDao.Get(db, userID)
	if err != nil {
		log.Log().Errorln(err)
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*30)
	}

	t := utils.NewToken()
	u.VerificationToken = &t

	if err := userDao.Update(db, u); err != nil {
		db.Rollback()

		log.Log().Errorln(err)
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*30)
	}

	if err := services.SendSignUpVerificationEmail(u.Name, u.Email, u.ID, *u.VerificationToken); err != nil {
		db.Rollback()

		log.Log().Errorln(err)
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*30)
	}

	if err := db.Commit().Error; err != nil {
		return tasks.NewErrRetryTaskLater(err.Error(), time.Second*30)
	}

	return nil
}
