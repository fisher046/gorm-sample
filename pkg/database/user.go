package database

import (
	"github.com/jinzhu/gorm"
)

// User table definition
type User struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(100)"`
	Email *string `gorm:"type:varchar(100);unique_index;not null"`
	Phone string  `gorm:"type:varchar(100)"`
	Extra string  `gorm:"type:varchar(100)"`
}

func init() {
	err := GetDB().AutoMigrate(&User{}).Error
	if err != nil {
		panic(err)
	}
}

// CreateUser creates one record in user table
func CreateUser(user *User) (id uint, err error) {
	err = GetDB().Create(user).Error
	if err != nil {
		return
	}
	return user.ID, nil
}

// GetUsers returns all users, temporarily do not support pagination
func GetUsers() (users []User, err error) {
	users = make([]User, 0)
	err = GetDB().Find(&users).Error
	return
}

// GetUser returns the info of one user by ID
func GetUser(id uint) (user *User, err error) {
	user = &User{Model: gorm.Model{ID: id}}
	err = GetDB().Take(user).Error
	if err != nil {
		user = nil
	}
	return
}

// DeleteUser deletes user record by ID
func DeleteUser(id uint) error {
	user := &User{Model: gorm.Model{ID: id}}
	err := GetDB().Delete(user).Error
	return err
}
