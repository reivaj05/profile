package auth

import (
	"fmt"

	"github.com/reivaj05/profile/db"
	"github.com/reivaj05/profile/users"
	"golang.org/x/crypto/bcrypt"

	"github.com/reivaj05/GoJSON"
)

func signup(data *GoJSON.JSONWrapper) error {
	fmt.Println("TODO: Implement signup")
	return nil
}

func resetPassword() error {
	fmt.Println("TODO: Implement resetPassword")
	return nil
}

func login(data *GoJSON.JSONWrapper) error {
	var user users.User
	email, _ := data.GetStringFromPath("email")
	password, _ := data.GetStringFromPath("password")
	err := db.DB.Client.First(&user, "email = ?", email).Error
	if err != nil {
		return err
	}
	return validatePassword(password, user.Password)
}

func validatePassword(password, passwordDB string) error {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8); err != nil {
		return err
	} else {
		if string(hashedPassword) == passwordDB {
			return nil
		}
		return fmt.Errorf("Wrong password")
	}
}

func logout() error {
	fmt.Println("TODO: Implement logout")
	return nil
}
