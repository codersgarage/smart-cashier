package migration

import (
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/core"
	"github.com/codersgarage/smart-cashier/log"
	"github.com/codersgarage/smart-cashier/models"
	"github.com/spf13/cobra"
)

var MigDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "drop drops database tables",
	Run:   drop,
}

func drop(cmd *cobra.Command, args []string) {
	tx := app.DB().Begin()

	var tables []core.Table
	tables = append(tables, &models.Entry{}, &models.Category{})
	tables = append(tables, &models.Diary{}, &models.Session{}, &models.User{})

	for _, t := range tables {
		if err := tx.DropTableIfExists(t).Error; err != nil {
			tx.Rollback()
			return
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Log().Errorln(err)
	}

	log.Log().Infoln("Migration drop completed")
}
