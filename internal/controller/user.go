package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "yx-blog/api/v1"
	"yx-blog/internal/consts"
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
		err = gerror.NewCode(gcode.New(consts.ErrUserExist, consts.ErrUserExistMsg, nil))
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

// DeleteUser 删除用户
func (c *cUser) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	//暂时关闭此接口
	err = gerror.NewCode(gcode.New(10007, "暂时关闭此接口", nil))

	//接收参数
	userId := req.UserId

	isHave, err := user.CheckUserById(ctx, userId)

	if !isHave {
		err = gerror.NewCode(gcode.New(consts.ErrUserNotExist, consts.ErrUserNotExistMsg, nil))
		return res, err
	}

	//删除用户
	err = user.Delete(ctx, userId)
	res = &v1.DeleteUserRes{
		UserId: userId,
	}

	return res, err
}
