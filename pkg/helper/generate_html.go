package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path"
)

//GenerateVerification will
func GenerateVerification(name, link string) (result string, err error) {
	pwd, _ := os.Getwd()
	filepath := path.Join(pwd, "/web/template/send_verification_email.html")
	template, err := template.ParseFiles(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var data = map[string]interface{}{
		"name": name,
		"link": link,
	}
	var temp bytes.Buffer
	err = template.Execute(&temp, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result = temp.String()
	return
}

//GenerateSendOTP will
func GenerateSendOTP(name, otp string) (result string, err error) {
	pwd, _ := os.Getwd()
	filepath := path.Join(pwd, "/web/template/send_otp_code.html")
	template, err := template.ParseFiles(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var data = map[string]interface{}{
		"name": name,
		"otp":  otp,
	}
	var temp bytes.Buffer
	err = template.Execute(&temp, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result = temp.String()
	return
}

//GenerateResetPassword will
func GenerateResetPassword(name, otp string) (result string, err error) {
	pwd, _ := os.Getwd()
	filepath := path.Join(pwd, "/web/template/send_reset_password.html")
	template, err := template.ParseFiles(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var data = map[string]interface{}{
		"name": name,
		"otp":  otp,
	}
	var temp bytes.Buffer
	err = template.Execute(&temp, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result = temp.String()
	return
}
