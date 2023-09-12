package authorization

import (
	"errors"
	"goapp/dao"
)

var (
	userCredentialKey = "user-credential"
)

func IsUserNameExist(username string) bool {
	isExist, _ := dao.GetRedisClient().HExists(userCredentialKey, username).Result()
	return isExist
}

func CreateUserCredential(username string, password string) (bool, error) {
	isExist := IsUserNameExist(username)
	if isExist {
		return false, errors.New("username has already existed")
	}
	success, err := dao.GetRedisClient().HSet(userCredentialKey, username, password).Result()
	return success, err
}

func DeleteUserCredential(username string, password string) bool {
	_, err := dao.GetRedisClient().HDel(userCredentialKey, username).Result()
	return err == nil
}
