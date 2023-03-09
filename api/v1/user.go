package v1

import "github.com/gogf/gf/v2/frame/g"

type AddUserReq struct {
	g.Meta   `path:"/user/add" tags:"用户" method:"post" summary:"添加用户"`
	UserName string `json:"user_name" example:"admin" format:"string" description:"用户名"`
	Password string `json:"password" example:"123456" format:"string" description:"密码"`
}

type AddUserRes struct {
	UserId     int64  `json:"user_id" example:"1" format:"int64" description:"用户ID"`
	UserName   string `json:"user_name" example:"admin" format:"string" description:"用户名"`
	IsBanned   bool   `json:"is_banned" example:"false" format:"bool" description:"是否禁用"`
	CreateTime string `json:"create_time" example:"2021-01-01 00:00:00" format:"string" description:"创建时间"`
}

type DeleteUserReq struct {
	g.Meta `path:"/user/delete" tags:"用户" method:"post" summary:"删除用户"`
	UserId int64 `json:"user_id" example:"1" format:"int64" description:"用户ID"`
}

type DeleteUserRes struct {
	UserId int64 `json:"user_id" example:"1" format:"int64" description:"用户ID"`
}
