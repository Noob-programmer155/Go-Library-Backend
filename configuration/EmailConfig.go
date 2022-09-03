package configuration

import (
	"amrDev/libraryBackend.com/dtos"
	"amrDev/libraryBackend.com/models"
	"amrDev/libraryBackend.com/repository"
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
)

func sendEmailValidation(w http.ResponseWriter, username string, email string) bool {
	from := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	to := []string{email}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	template, err := template.ParseFiles("../template/EmailTemplate.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Email Validation \n%s\n\n", mimeHeaders)))

	template.Execute(&body, struct {
		Name  string
		Token string
	}{
		Name:  username,
		Token: base64.StdEncoding.EncodeToString([]byte(email)),
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	return true
}

func emailValidation(w http.ResponseWriter, token string, res *http.Response) *dtos.UserInfoDTO {
	email, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return &dtos.UserInfoDTO{}
	}
	if isChecked, _ := repository.CheckByEmail(string(email)); isChecked {
		user, err1 := repository.FindIdNameEmailByEmail(string(email))
		user.Role = models.USER
		err2 := repository.Update(user)
		if err1 != nil || err2 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return &dtos.UserInfoDTO{}
		}
		if token, isCreated := CreateToken(user.Name, user.Email, models.USER); isCreated {
			if token != "" {
				Bind(w, token)
			}
			return &dtos.UserInfoDTO{
				Id:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			}
		}
	}
	return &dtos.UserInfoDTO{}
}
