package database

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func resetUserPasswordTable() {
	err := GetDB().DropTableIfExists(&UserPassword{}).
		AutoMigrate(&UserPassword{}).Error
	if err != nil {
		panic(err)
	}
}

func TestCreateUserPassword(t *testing.T) {
	resetUserPasswordTable()

	testCases := []struct {
		userID string
	}{
		{userID: "test"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			id, err := CreateUserPassword(&UserPassword{UserID: tc.userID})
			if err != nil {
				t.Fatal(err)
			}
			userPwd := &UserPassword{Model: gorm.Model{ID: id}}
			GetDB().Take(userPwd)
			if userPwd.UserID != tc.userID {
				t.Fatalf("expected: %v, actual: %v", tc.userID, userPwd.UserID)
			}
		})
	}
}
