package v1

import "github.com/gogf/gf/v2/frame/g"

type AddUserReq struct {
	g.Meta   `path:"/user/add" tags:"用户" method:"post" summary:"添加用户"`
	UserName string `json:"user_name" example:"admin" format:"string" description:"用户名"`
	Password string `json:"password" example:"123456" format:"string" description:"密码"`
	MyKey    string `json:"my_key" example:"2d8f7e20fb8d2784a855a50dd71bd5a2" format:"string" description:"密钥"`
}

type AddUserRes struct {
	UserId     int64  `json:"user_id" example:"1" format:"int64" description:"用户ID"`
	UserName   string `json:"user_name" example:"admin" format:"string" description:"用户名"`
	IsBanned   bool   `json:"is_banned" example:"false" format:"bool" description:"是否禁用"`
	CreateTime string `json:"create_time" example:"2021-01-01 00:00:00" format:"string" description:"创建时间"`
}

type DeleteUserReq struct {
	g.Meta `path:"/user/delete" tags:"用户" method:"post" summary:"删除用户"`
	UserId int64  `json:"user_id" example:"1" format:"int64" description:"用户ID"`
	MyKey  string `json:"my_key" example:"2d8f7e20fb8d2784a855a50dd71bd5a2" format:"string" description:"密钥"`
}

type DeleteUserRes struct {
	UserId int64 `json:"user_id" example:"1" format:"int64" description:"用户ID"`
}

// UpdatePasswordReq 修改密码
type UpdatePasswordReq struct {
	g.Meta      `path:"/user/update/password" tags:"用户" method:"post" summary:"修改密码"`
	UserId      int64  `json:"user_id" example:"1" format:"int64" description:"用户ID"`
	OldPassword string `json:"old_password" example:"123456" format:"string" description:"旧密码"`
	NewPassword string `json:"new_password" example:"123456" format:"string" description:"新密码"`
	MyKey       string `json:"my_key" example:"2d8f7e20fb8d2784a855a50dd71bd5a2" format:"string" description:"密钥"`
}

type UpdatePasswordRes struct {
	UserId int64 `json:"user_id" example:"1" format:"int64" description:"用户ID"`
}

// UserListReq 查询用户列表
type UserListReq struct {
	g.Meta `path:"/user/list" tags:"用户" method:"post" summary:"查询用户列表"`
	//分页参数

	Page  int    `json:"page" example:"1" format:"int" description:"页码"`
	Limit int    `json:"limit" example:"10" format:"int" description:"每页条数"`
	MyKey string `json:"my_key" example:"2d8f7e20fb8d2784a855a50dd71bd5a2" format:"string" description:"密钥"`
}

type UserListRes struct {
	List []UserInfo `json:"list" example:"[]" format:"[]" description:"用户列表"`
}

type UserInfo struct {
	UserId     int64  `json:"user_id" example:"1" format:"int64" description:"用户ID"`
	UserName   string `json:"user_name" example:"admin" format:"string" description:"用户名"`
	Email      string `json:"email" example:"email" format:"string" description:"邮箱"`
	IsBanned   bool   `json:"is_banned" example:"false" format:"bool" description:"是否禁用"`
	CreateTime string `json:"create_time" example:"2021-01-01 00:00:00" format:"string" description:"创建时间"`
	UpdateTime string `json:"update_time" example:"2021-01-01 00:00:00" format:"string" description:"更新时间"`
}

// UserInfoReq 获取单个用户信息
type UserInfoReq struct {
	g.Meta `path:"/user/info" tags:"用户" method:"post" summary:"获取单个用户信息"`
	UserId int64  `json:"user_id" example:"1" format:"int64" description:"用户ID"`
	MyKey  string `json:"my_key" example:"2d8f7e20fb8d2784a855a50dd71bd5a2" format:"string" description:"密钥"`
}

type UserInfoRes struct {
	UserId     int64  `json:"user_id" example:"1" format:"int64" description:"用户ID"`
	UserName   string `json:"user_name" example:"admin" format:"string" description:"用户名"`
	Email      string `json:"email" example:"email" format:"string" description:"邮箱"`
	IsBanned   bool   `json:"is_banned" example:"false" format:"bool" description:"是否禁用"`
	CreateTime string `json:"create_time" example:"2021-01-01 00:00:00" format:"string" description:"创建时间"`
	UpdateTime string `json:"update_time" example:"2021-01-01 00:00:00" format:"string" description:"更新时间"`
}
