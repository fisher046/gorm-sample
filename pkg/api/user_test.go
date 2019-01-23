package api

import (
	"testing"

	"github.com/fisher046/gorm-sample/pkg/database"
)

func resetTables() {
	err := database.GetDB().DropTableIfExists(&database.User{}).
		AutoMigrate(&database.User{}).Error
	if err != nil {
		panic(err)
	}

	err = database.GetDB().DropTableIfExists(&database.UserPassword{}).
		AutoMigrate(&database.UserPassword{}).Error
	if err != nil {
		panic(err)
	}
}

func TestDoCreateUser(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		email    string
	}{
		{name: "test", password: "pass", email: "email"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			resetTables()
			user := &User{
				Name:     tc.name,
				Password: tc.password,
				Email:    tc.email,
			}
			id, err := doCreateUser(user)
			if err != nil {
				t.Fatal(err)
			}

			if id != 1 {
				t.Fatalf("expected: 1, actual: %v", id)
			}

			userInDB, err := database.GetUser(id)
			if err != nil {
				t.Fatal(err)
			}
			if *userInDB.Email != tc.email {
				t.Fatalf("expected: %v, actual: %v", tc.email, *userInDB.Email)
			}
		})
	}
}
