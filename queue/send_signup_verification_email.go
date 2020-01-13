package queue

import (
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/codersgarage/smart-cashier/machinery"
	tasks2 "github.com/codersgarage/smart-cashier/tasks"
)

func SendSignUpVerificationEmail(userID string) error {
	sig := &tasks.Signature{
		Name: tasks2.SendSignUpVerificationEmailTaskName,
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: userID,
				Name:  "userID",
			},
		},
	}
	_, err := machinery.RabbitMQConnection().SendTask(sig)
	if err != nil {
		return err
	}
	return nil
}
