package database

import (
	"github.com/jinzhu/gorm"
)

// UserPassword table definition
type UserPassword struct {
	gorm.Model
	UserID uint
	Hash   []byte
}

func init() {
	err := GetDB().AutoMigrate(&UserPassword{}).Error
	if err != nil {
		panic(err)
	}
}

// CreateUserPassword creates a user-password relationship
func CreateUserPassword(userPwd *UserPassword) (id uint, err error) {
	err = GetDB().Create(userPwd).Error
	if err != nil {
		return
	}
	return userPwd.ID, nil
}
