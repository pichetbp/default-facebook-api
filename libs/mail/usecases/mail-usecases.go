package usecases

import (
	"default-repo/libs/mail/models"
	"errors"
	"strings"

	"github.com/spf13/viper"
	mailLib "github.com/xhit/go-simple-mail/v2"
)

type MailInterface interface {
}

type mailIns struct {
	server *mailLib.SMTPClient
}

func NewMail() MailInterface {

	server := mailLib.NewSMTPClient()
	server.Host = viper.GetString("sendMail.host")
	server.Port = viper.GetInt("sendMail.port")
	server.Username = viper.GetString("sendMail.username")
	server.Password = viper.GetString("sendMail.password")

	server.Encryption = mailLib.EncryptionSTARTTLS
	server.Authentication = mailLib.AuthLogin

	// Connect server smtp
	smtpClient, errConnectSmtp := server.Connect()
	if errConnectSmtp != nil {
		return errors.New("cannot connect to SMTP server")
	}

	return &mailIns{
		server: smtpClient,
	}
}

func (m *mailIns) SendMail(body models.BodyMailRequest) error {
	err := body.ValidateMail()
	if err != nil {
		return err
	}

	// Create email
	reqSending := models.BodyMailRequest{
		To:       body.To,
		Subject:  body.Subject,
		Body:     body.Body,
		TypeBody: body.TypeBody,
		Cc:       body.Cc,
		Attach:   body.Attach,
		Priority: body.Priority,
	}
	errSending := m.sending(reqSending)
	if errSending != nil {
		return errors.New("cannot send email")
	}

	return nil
}

func (m *mailIns) sending(req models.BodyMailRequest) error {
	email := mailLib.NewMSG()
	email.SetFrom(viper.GetString("sendMail.username")).
		SetSubject(req.Subject)

	// Add To
	if req.To != "" {
		toList := strings.Split(req.To, ",")
		if len(toList) > 0 {
			for _, to := range toList {
				email.AddTo(to)
			}
		}
	}

	// Add CC
	if req.Cc != "" {
		ccList := strings.Split(req.Cc, ",")
		if len(ccList) > 0 {
			for _, cc := range ccList {
				email.AddCc(cc)
			}
		}
	}

	// Check type body message
	switch strings.ToUpper(req.TypeBody) {
	case "HTML":
		email.SetBody(mailLib.TextHTML, req.Body)
	case "TEXT":
		email.SetBody(mailLib.TextPlain, req.Body)
	}

	// Add attach file
	if len(req.Attach) > 0 {
		for i := range req.Attach {
			email.Attach(&mailLib.File{FilePath: req.Attach[i]})
		}
	}

	// Set Priority email
	if req.Priority != "" {
		switch strings.ToUpper(req.Priority) {
		case "HIGH":
			email.SetPriority(mailLib.PriorityHigh)
		case "LOW":
			email.SetPriority(mailLib.PriorityLow)
		}
	}

	// Always check error after send
	if email.Error != nil {
		return errors.New("Cannot Send Email")
	}

	// Send email
	// errSendMail := email.Send(m.server)
	// if errSendMail != nil {
	// 	return errors.New("Cannot Send Email")
	// }

	return nil
}
