package entity

import (
	"encoding/json"
	"kill8s/infrastructure/commons/util"
	"kill8s/infrastructure/constant"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (u *User) MD5Password() error {
	md5, err := util.MD5WithSalt(u.Password, constant.Md5Salt)
	if err != nil {
		return err
	}
	u.Password = md5
	return nil
}

func (u *User) String() (string, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
