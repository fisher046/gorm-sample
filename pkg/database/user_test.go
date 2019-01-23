package database

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func resetUserTable() {
	err := GetDB().DropTableIfExists(&User{}).AutoMigrate(&User{}).Error
	if err != nil {
		panic(err)
	}
}

func TestCreateUser(t *testing.T) {
	resetUserTable()

	testCases := []struct {
		name   string
		email  string
		errMsg string
	}{
		{errMsg: "NOT NULL constraint failed: users.email"},
		{name: "test name", email: "test email"},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			var id uint
			var err error
			if tc.email != "" {
				id, err = CreateUser(&User{Name: tc.name, Email: &tc.email})
			} else {
				id, err = CreateUser(&User{Name: tc.name})
			}

			if err != nil {
				if err.Error() != tc.errMsg {
					t.Fatalf("expected: %v, actual: %v", tc.errMsg, err.Error())
				}
				if id != 0 {
					t.Fatalf("expected: 0, actual: %v", id)
				}
			} else {
				if tc.errMsg != "" {
					t.Fatalf("expected: %v, actual: nil", tc.errMsg)
				}
				user := &User{Model: gorm.Model{ID: id}}
				GetDB().Take(user)
				if user.Name != tc.name {
					t.Fatalf("expected: %v, actual: %v", tc.name, user.Name)
				}
				if *user.Email != tc.email {
					t.Fatalf("expected: %v, actual: %v", tc.email, *user.Email)
				}
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	testCases := []struct {
		count int
	}{
		{0}, {2},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			resetUserTable()
			for i := 0; i < tc.count; i++ {
				email := "test" + string(i)
				GetDB().Create(&User{Email: &email})
			}

			users, err := GetUsers()
			if err != nil {
				t.Fatal(err)
			}

			if len(users) != tc.count {
				t.Error(users)
				t.Fatalf("expected: %v, actual: %v", tc.count, len(users))
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	testCases := []struct {
		users  []*User
		id     uint
		phone  string
		errMsg string
	}{
		{errMsg: ErrDBNotFound},
		{
			users:  []*User{&User{Phone: "123"}},
			id:     2,
			errMsg: ErrDBNotFound,
		},
		{
			users: []*User{&User{Phone: "456"}, &User{Phone: "789"}},
			id:    2,
			phone: "789",
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			resetUserTable()
			for _, v := range tc.users {
				email := "test" + v.Phone
				v.Email = &email
				GetDB().Create(v)
			}

			user, err := GetUser(tc.id)
			if err != nil {
				if err.Error() != tc.errMsg {
					t.Fatalf("expected: %v, actual: %v", tc.errMsg, err.Error())
				}
				if user != nil {
					t.Fatalf("expected: nil, actual: %v", *user)
				}
			} else {
				if tc.errMsg != "" {
					t.Fatalf("expected: %v, actual: nil", tc.errMsg)
				}
				if user.Phone != tc.phone {
					t.Fatalf("expected: %v, actual: %v", tc.phone, user.Phone)
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	testCases := []struct {
		users []*User
		id    uint
	}{
		{id: 2},
		{
			users: []*User{&User{Phone: "456"}, &User{Phone: "789"}},
			id:    2,
		},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			resetUserTable()
			for _, v := range tc.users {
				GetDB().Create(v)
			}

			err := DeleteUser(tc.id)
			if err != nil {
				t.Fatal(err)
			}

			err = GetDB().Take(&User{Model: gorm.Model{ID: tc.id}}).Error
			if err.Error() != ErrDBNotFound {
				t.Fatalf("expected: %v, actual: %v", ErrDBNotFound, err.Error())
			}
		})
	}
}
