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
	err = gerror.NewCode(gcode.New(consts.ErrCloseThisInterface, consts.ErrCloseThisInterfaceMsg, nil))
	return res, err

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

// UpdatePassword 修改用户密码
func (c *cUser) UpdatePassword(ctx context.Context, req *v1.UpdatePasswordReq) (res *v1.UpdatePasswordRes, err error) {
	//接收参数
	userId := req.UserId
	oldPassword := req.OldPassword
	newPassword := req.NewPassword

	//查询是否存在用户
	isHave, err := user.CheckUserById(ctx, userId)
	//判断是否存在用户
	if !isHave {
		err = gerror.NewCode(gcode.New(consts.ErrUserNotExist, consts.ErrUserNotExistMsg, nil))
		return res, err
	}

	//查询旧密码是否正确
	isRight, err := user.CheckPassword(ctx, userId, oldPassword)
	//判断是否正确
	if !isRight {
		err = gerror.NewCode(gcode.New(consts.ErrUserPasswordError, consts.ErrUserPasswordErrorMsg, nil))
		return res, err
	}

	//修改密码
	err = user.UpdatePassword(ctx, userId, newPassword)

	res = &v1.UpdatePasswordRes{
		UserId: userId,
	}

	return res, err
}

// UserList 用户列表
func (c *cUser) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	//接收参数
	page := req.Page
	limit := req.Limit
	//查询用户列表
	list, err := user.List(ctx, page, limit)
	//返回数据
	res = &v1.UserListRes{
		List: list,
	}
	return res, err
}
