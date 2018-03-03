package service

import (
	"strings"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
)

type UserService struct{}

func (this *UserService) GetUserByEmail(email string) (users *info.Users) {
	email = strings.ToLower(email)
	users = &info.Users{Email: email}
	_, err := db.MasterDB.Get(users)
	if err != nil {
		revel.INFO.Printf("GetUserByEmail Fatal : %s", err)
	}
	return users
}

func (this *UserService) UpdatePassword(id int, password string) bool {
	newPassword,hashErr := authService.HashPassword(password)
	if hashErr != nil {
		revel.INFO.Printf("Update hash password error: %s", hashErr)
		return false
	}
	user := info.Users{Password: newPassword}
	_, err := db.MasterDB.Id(id).Cols("password").Update(user)
	if err != nil {
		revel.INFO.Printf("Update by name option error: %s", err)
		return false
	}
	return true
}
