package auth

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"os"
	"time"

	"github.com/reivaj05/profile/db"
	"github.com/reivaj05/profile/users"
	"golang.org/x/crypto/bcrypt"

	"github.com/reivaj05/GoJSON"
)

func resetPassword(data *GoJSON.JSONWrapper) error {
	var user users.User
	email, _ := data.GetStringFromPath("email")
	err := db.DB.Client.First(&user, "email = ?", email).Error
	if err != nil {
		return err
	}
	return sendPassword(user, email)
}

func sendPassword(user users.User, email string) error {
	from := os.Getenv("GMAIL_ACCOUNT")
	pass := os.Getenv("GMAIL_PASSWORD")
	newPassword := createNewPassword()
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: Restore Password\n%s", from, email, newPassword)
	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{email}, []byte(msg))
	if err == nil {
		saveNewPassword(user, newPassword)
	}
	return err
}

func createNewPassword() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 15)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func saveNewPassword(user users.User, newPassword string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), 8)
	user.Password = string(hashedPassword)
	db.DB.Client.Save(user)
}

func login(data *GoJSON.JSONWrapper) (*users.User, error) {
	var user users.User
	email, _ := data.GetStringFromPath("email")
	password, _ := data.GetStringFromPath("password")
	err := db.DB.Client.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return &user, bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
