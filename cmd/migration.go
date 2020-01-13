package cmd

import (
	migration "github.com/codersgarage/smart-cashier/cmd/migrations"
	"github.com/spf13/cobra"
)

var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration migrates database schemas",
}

func init() {
	migrationCmd.AddCommand(migration.MigDropCmd)
	migrationCmd.AddCommand(migration.MigAutoCmd)
}
