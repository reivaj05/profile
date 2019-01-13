package users

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"

	"github.com/reivaj05/profile/db"

	"github.com/jinzhu/gorm"
	"github.com/reivaj05/GoJSON"
)

type User struct {
	gorm.Model
	Email     string `gorm:"type:varchar(50);not_null;unique" json:"email"`
	Password  string `gorm:"type:varchar(150);not_null"`
	FirstName string `gorm:"type:varchar(50)" json:"firstName"`
	LastName  string `gorm:"type:varchar(50)" json:"lastName"`
	Phone     string `gorm:"type:varchar(50)" json:"phone"`
	Address   string `gorm:"type:text;" json:"address"`
}

func newUser(data *GoJSON.JSONWrapper) (*User, error) {
	email, _ := data.GetStringFromPath("email")
	password, _ := data.GetStringFromPath("password")
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return nil, err
	}
	user := &User{Email: email, Password: string(hashedPassword)}
	err = db.DB.Client.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getUser(id int) (*User, error) {
	var user User
	err := db.DB.Client.First(&user, "ID = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (user *User) update(data *GoJSON.JSONWrapper) error {
	user.Email, _ = data.GetStringFromPath("email")
	user.FirstName, _ = data.GetStringFromPath("firstName")
	user.LastName, _ = data.GetStringFromPath("lastName")
	user.Address, _ = data.GetStringFromPath("address")
	user.Phone, _ = data.GetStringFromPath("phone")
	return db.DB.Client.Save(user).Error
}

func (user *User) ToJSON() string {
	data, _ := json.Marshal(user)
	return string(data)
}
