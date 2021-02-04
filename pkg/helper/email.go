package helper

import (
	"github.com/spf13/viper"
	log "go.uber.org/zap"
	"gopkg.in/gomail.v2"
)

//SendMail will
func SendMail(receiver, subject, body, emailType string) {

	dialer := gomail.NewDialer(
		viper.GetString(`smtp.host`),
		viper.GetInt(`smtp.port`),
		viper.GetString(`smtp.username`),
		viper.GetString(`smtp.password`),
	)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", viper.GetString(`smtp.email`))
	mailer.SetHeader("To", receiver)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.S().Errorf("error send email : %s | %s ", emailType, err.Error())
		return
	}

	log.S().Infof("success send email : %s | %s ", emailType, receiver)
}
