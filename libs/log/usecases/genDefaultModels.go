package usecases

import (
	"default-repo/helpers"
	"default-repo/libs/log/models"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenDefaultModels() models.LogDefault {
	return models.LogDefault{
		Account:       "default",
		FunctionName:  "default",
		Level:         logrus.DebugLevel.String(),
		TimeStamp:     time.Now(),
		TransactionId: uuid.New().String(),
		FileLocation:  helpers.SetFileLocation(),
		Type:          "default",
	}
}
