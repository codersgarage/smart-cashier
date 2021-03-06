package machinery

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	cfg "github.com/codersgarage/smart-cashier/config"
	"github.com/codersgarage/smart-cashier/log"
	"github.com/codersgarage/smart-cashier/tasks"
	"os"
)

var machineryServer *machinery.Server
var msErr error

func NewRabbitMQConnection() error {
	machineryServer, msErr = machinery.NewServer(&config.Config{
		Broker:        cfg.RabbitMQ().Broker,
		DefaultQueue:  cfg.RabbitMQ().DefaultQueue,
		ResultBackend: cfg.RabbitMQ().ResultBackend,
		AMQP: &config.AMQPConfig{
			ExchangeType:  cfg.RabbitMQ().AMQP.ExchangeType,
			Exchange:      cfg.RabbitMQ().AMQP.Exchange,
			BindingKey:    cfg.RabbitMQ().AMQP.BindingKey,
			PrefetchCount: cfg.RabbitMQ().AMQP.PrefetchCount,
		},
		ResultsExpireIn: 3600,
	})
	if msErr != nil {
		return msErr
	}
	return nil
}

func RabbitMQConnection() *machinery.Server {
	return machineryServer
}

func RegisterRabbitMQTasks() error {
	if err := machineryServer.RegisterTask(tasks.SendSignUpVerificationEmailTaskName, tasks.SendSignUpVerificationEmailFn); err != nil {
		return err
	}
	return nil
}

var worker *machinery.Worker
var err error

func RunRabbitMQWorker() {
	cnf := cfg.RabbitMQ().Worker
	worker = RabbitMQConnection().NewWorker(cnf.Name, cnf.Count)
	err = worker.Launch()
	if err != nil {
		log.Log().Errorln("Couldn't launch worker", err)
		os.Exit(-1)
	}
}
