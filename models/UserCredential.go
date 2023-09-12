package models

import "goapp/dao/authorization"

type UserCredential struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (uc *UserCredential) IsExist() bool {
	return authorization.IsUserNameExist(uc.UserName)
}

func (uc *UserCredential) Create() (bool, error) {
	return authorization.CreateUserCredential(uc.UserName, uc.Password)
}

func (uc *UserCredential) Delete() bool {
	return authorization.DeleteUserCredential(uc.UserName, uc.Password)
}
