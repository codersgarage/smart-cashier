package migration

import (
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/core"
	"github.com/codersgarage/smart-cashier/log"
	"github.com/codersgarage/smart-cashier/models"
	"github.com/spf13/cobra"
	"strings"
)

var MigAutoCmd = &cobra.Command{
	Use:   "auto",
	Short: "auto alter database tables if required",
	Run:   auto,
}

func auto(cmd *cobra.Command, args []string) {
	tx := app.DB().Begin()

	var tables []core.Table
	tables = append(tables, &models.User{}, &models.Session{})
	tables = append(tables, &models.Diary{}, &models.Entry{}, &models.Category{})

	for _, t := range tables {
		if err := tx.AutoMigrate(t).Error; err != nil {
			tx.Rollback()
			log.Log().Errorln(err)
			return
		}
	}

	var tForeignKeys []core.Model
	tForeignKeys = append(tForeignKeys, &models.Session{}, &models.Diary{}, &models.Entry{}, &models.Category{})

	for _, t := range tForeignKeys {
		for _, fks := range t.ForeignKeys() {
			fk := strings.Split(fks, ";")
			if err := tx.Model(t).AddForeignKey(fk[0], fk[1], fk[2], fk[3]).Error; err != nil {
				tx.Rollback()
				log.Log().Errorln(err)
				return
			}
		}
	}

	var views []core.View
	//views = append(views, &models.StoreUserProfile{})

	for _, v := range views {
		if err := v.CreateView(tx); err != nil {
			tx.Rollback()
			log.Log().Errorln(err)
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Log().Errorln(err)
		return
	}

	log.Log().Infoln("Migration auto completed")
}
