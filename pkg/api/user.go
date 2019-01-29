package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/crypto/bcrypt"

	"github.com/fisher046/gorm-sample/pkg/database"
)

// User structure
type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Phone string `json:"phone"`
}

// UserPassword is the structure of user info with password
type UserPassword struct {
	User
	Password string `json:"password" binding:"required"`
}

// CreateUser is the implementation of /users
func CreateUser(c *gin.Context) {
	var usrPwd UserPassword

	err := c.ShouldBindJSON(&usrPwd)
	if err != nil {
		glog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := doCreateUser(&usrPwd)
	if err != nil {
		glog.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})

	return
}

func doCreateUser(usrPwd *UserPassword) (id uint, err error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(usrPwd.Password), bcrypt.DefaultCost,
	)
	if err != nil {
		return
	}

	id, err = database.CreateUser(
		&database.User{
			Name:  usrPwd.Name,
			Email: &usrPwd.Email,
			Phone: usrPwd.Phone,
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
