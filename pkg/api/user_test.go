package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

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

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		email    string
		code     int
	}{
		{code: http.StatusBadRequest},
		{
			name: "test", password: "pass", email: "email",
			code: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			resetTables()
			r := gin.Default()
			r.POST("/user", CreateUser)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				"POST", "/user",
				strings.NewReader(
					fmt.Sprintf(
						"{\"Name\":\"%v\",\"Email\":\"%v\",\"Password\":\"%v\"}",
						tc.name, tc.email, tc.password,
					),
				),
			)
			r.ServeHTTP(w, req)
			if w.Code != tc.code {
				t.Errorf("response body: %v", w.Body.String())
				t.Fatalf("expected: %v, actual: %v", tc.code, w.Code)

			}
		})
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
			usrPwd := &UserPassword{
				User: User{
					Name:  tc.name,
					Email: tc.email,
				},
				Password: tc.password,
			}
			id, err := doCreateUser(usrPwd)
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
