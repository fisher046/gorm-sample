package api

import (
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"

	"github.com/fisher046/gorm-sample/pkg/database"
)

// User structure
type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"`
}

func doCreateUser(user *User) (id uint, err error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password), bcrypt.DefaultCost,
	)
	if err != nil {
		return
	}

	id, err = database.CreateUser(
		&database.User{
			Name:  user.Name,
			Email: &user.Email,
			Phone: user.Phone,
		},
	)
	if err != nil {
		return 0, err
	}

	_, err = database.CreateUserPassword(
		&database.UserPassword{
			UserID: id,
			Hash:   hash,
		},
	)
	if err != nil {
		e := database.DeleteUser(id)
		if e != nil {
			glog.Error(err)
		}
		return 0, err
	}

	return
}
