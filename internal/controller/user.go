package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "yx-blog/api/v1"
	"yx-blog/internal/logic/user"
)

var (
	User = cUser{}
)

type cUser struct {
}

// AddUser 添加用户
func (c *cUser) AddUser(ctx context.Context, req *v1.AddUserReq) (res *v1.AddUserRes, err error) {
	//接收参数
	userName := req.UserName
	//查询是否存在用户
	isHave, err := user.Check(ctx, userName)
	//判断是否存在用户
	if isHave {
		err = gerror.NewCode(gcode.New(10001, "用户已存在", nil))
		return res, err
	}

	password := req.Password
	id, createTime := user.Add(ctx, userName, password)
	//返回数据
	res = &v1.AddUserRes{
		UserId:     id,
		UserName:   userName,
		IsBanned:   false,
		CreateTime: createTime,
	}
	return res, err
}
