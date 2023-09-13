package user

import "kill8s/domain/entity"

type RegisterReq struct {
	UserName string `json:"user_name" form:"name"`
	Password string `json:"password" form:"password"`
}

type RegisterResp struct {
	User *entity.User
}
