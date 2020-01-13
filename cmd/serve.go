package cmd

import (
	"github.com/codersgarage/smart-cashier/app"
	"github.com/codersgarage/smart-cashier/log"
	"github.com/codersgarage/smart-cashier/machinery"
	"github.com/codersgarage/smart-cashier/server"
	"github.com/spf13/cobra"
	"os"
)

var serveCmd = &cobra.Command{
	Use:    "serve",
	Short:  "Serve starts http server",
	PreRun: preServe,
	Run:    serve,
}

func preServe(cmd *cobra.Command, args []string) {
	if err := app.ConnectMinio(); err != nil {
		log.Log().Errorln("Failed to connect to minio : ", err)
		os.Exit(-1)
	}

	if err := machinery.NewRabbitMQConnection(); err != nil {
		log.Log().Errorln("Failed to connect to rabbitmq : ", err)
		os.Exit(-1)
	}
	go machinery.RunRabbitMQWorker()

	if err := machinery.RegisterRabbitMQTasks(); err != nil {
		log.Log().Errorln("Failed to register rabbitmq tasks : ", err)
		os.Exit(-1)
	}
}

func serve(cmd *cobra.Command, args []string) {
	server.StartServer()
}
