package models

import (
	"errors"

	"github.com/spf13/viper"
)

type BodyMailRequest struct {
	To       string
	Subject  string
	Body     string
	TypeBody string
	Cc       string
	Attach   []string
	Priority string
}

func (body *BodyMailRequest) ValidateMail() error {

	switch {
	case body.To == "":
		return errors.New("to is empty")
	case body.Subject == "":
		return errors.New("subject is empty")
	case body.Body == "":
		return errors.New("body is empty")
	case body.TypeBody == "":
		return errors.New("typeBody is empty")
	case viper.GetString("sendMail.host") == "":
		return errors.New("unable to read sendMail.host from configuration file")
	case viper.GetInt("sendMail.port") == 0:
		return errors.New("unable to read sendMail.port from configuration file")
	case viper.GetString("sendMail.username") == "":
		return errors.New("unable to read sendMail.username from configuration file")
	case viper.GetString("sendMail.password") == "":
		return errors.New("unable to read sendMail.password from configuration file")
	}
	return nil
}
