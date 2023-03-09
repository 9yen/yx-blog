package user

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

var (
	Db = g.Model("users")
)

// Add 添加用户
func Add(ctx context.Context, userName string, password string) (id int64, createTime string) {
	//随机生成一个盐
	salt := grand.S(6)
	//密码加盐
	password = password + salt
	//md5加密
	password = gmd5.MustEncryptString(password)

	//获取当前时间
	createTime = gtime.Now().String()

	//插入数据
	result, err := Db.Insert(g.Map{
		"user_name":   userName,
		"password":    password,
		"salt":        salt,
		"create_time": createTime,
		"update_time": createTime,
		"is_banned":   0,
	})
	//把result转换为int64
	id, err = result.LastInsertId()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)

	}
	return
}

// Check 查询是否存在用户
func Check(ctx context.Context, userName string) (isHave bool, err error) {
	isHave = false
	//查询数据
	result, err := Db.Where("username", userName).One()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	//判断是否有数据
	if result != nil {
		isHave = true
	}
	return
}
