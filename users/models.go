package users

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(50);not_null"`
	Email       string `gorm:"type:varchar(50);not_null"`
	FirstName   string `gorm:"type:varchar(50)"`
	LastName    string `gorm:"type:varchar(50)"`
	Address     string `gorm:"type:text;"`
	Translation string `gorm:"type:text;not_null"`
}

func newUser() (*User, error) {
	fmt.Prinln("TODO: Implement create user")
	return &User{}, nil
}

func (user *User) update() error {
	fmt.Prinln("TODO: Implement update user")
	return nil
}

func (user *User) toJSON() string {
	fmt.Prinln("TODO: Implement user to json")
	return ""
}
