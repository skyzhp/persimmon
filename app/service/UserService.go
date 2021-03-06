package service

import (
	"strings"
	"github.com/cong5/persimmon/app/info"
	"github.com/cong5/persimmon/app/db"
	"github.com/revel/revel"
)

type UserService struct{}

func (this *UserService) GetUserByUid(uid int) (*info.Users, error) {
	users := &info.Users{Id: uid}
	_, err := db.MasterDB.Get(users)
	if err != nil {
		revel.INFO.Printf("GetUserByUid failed : %s", err)
		return nil, err
	}

	return users, nil
}

func (this *UserService) GetUserByEmail(email string) (*info.Users, error) {
	email = strings.ToLower(email)
	users := &info.Users{Email: email}
	_, err := db.MasterDB.Get(users)
	if err != nil {
		revel.INFO.Printf("GetUserByEmail failed : %s", err)
		return nil, err
	}

	return users, nil
}

func (this *UserService) UpdatePassword(id int, password string) (bool, error) {
	newPassword, hashErr := authService.HashPassword(password)
	if hashErr != nil {
		//revel.INFO.Printf("Update hash password failed: %s", hashErr)
		return false, hashErr
	}

	user := info.Users{Password: newPassword}
	_, err := db.MasterDB.Id(id).Cols("password").Update(user)
	if err != nil {
		//revel.INFO.Printf("Update password failed: %s", err)
		return false, hashErr
	}

	return true, nil
}

func (this *UserService) Table(tableName string) string {
	return db.MasterDB.TableMapper.Obj2Table(tableName)
}
