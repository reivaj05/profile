package users

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/reivaj05/GoJSON"
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

func newUser(data *GoJSON.JSONWrapper) (*User, error) {
	fmt.Println("TODO: Implement create user")
	return &User{}, nil
}

func getUser(id int) (*User, error) {
	fmt.Println("TODO: Implement get user")
	return &User{}, nil
}

func (user *User) update(data *GoJSON.JSONWrapper) error {
	fmt.Println("TODO: Implement update user")
	return nil
}

func (user *User) toJSON() string {
	fmt.Println("TODO: Implement user to json")
	return ""
}
