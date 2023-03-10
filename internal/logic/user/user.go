package user

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	v1 "yx-blog/api/v1"
)

// Add 添加用户
func Add(ctx context.Context, userName string, password string) (id int64, createTime string) {
	db := g.Model("users")
	//随机生成一个盐
	salt := grand.S(6)
	//密码加盐
	password = password + salt
	//md5加密
	password = gmd5.MustEncryptString(password)

	//获取当前时间
	createTime = gtime.Now().String()

	//插入数据
	result, err := db.Insert(g.Map{
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

// Delete 删除用户
func Delete(ctx context.Context, userId int64) (err error) {
	db := g.Model("users")
	//删除数据
	_, err = db.Where("id", userId).Delete()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	return
}

// UpdatePassword 修改用户密码
func UpdatePassword(ctx context.Context, userId int64, password string) (err error) {
	db := g.Model("users")
	//随机生成一个盐
	salt := grand.S(6)
	//密码加盐
	password = password + salt
	//md5加密
	password = gmd5.MustEncryptString(password)
	//获取当前时间
	updateTime := gtime.Now().String()
	//修改数据
	_, err = db.Data(g.Map{
		"password":    password,
		"salt":        salt,
		"update_time": updateTime,
	}).Where("id", userId).Update()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	return
}

// CheckPassword 验证密码是否正确
func CheckPassword(ctx context.Context, userId int64, password string) (isRight bool, err error) {
	db := g.Model("users")
	isRight = false
	//查询数据
	result, err := db.Where("id", userId).One()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	//判断是否有数据
	if result != nil {
		//获取盐
		salt := result["salt"]
		//密码加盐
		password = password + salt.String()
		//md5加密
		password = gmd5.MustEncryptString(password)
		//判断密码是否正确
		if password == result["password"].String() {
			isRight = true
		}
	}
	return
}

// Check 根据名字查询是否存在用户
func Check(ctx context.Context, userName string) (isHave bool, err error) {
	db := g.Model("users")
	isHave = false
	//查询数据
	result, err := db.Where("username", userName).One()

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

// CheckUserById 根据id查询用户
func CheckUserById(ctx context.Context, userId int64) (isHave bool, err error) {
	db := g.Model("users")
	//查询数据
	result, err := db.Where("id", userId).One()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}

	//判断是否有数据
	if result != nil {
		isHave = true
	} else {
		isHave = false
	}
	return
}

// List 用户列表
func List(ctx context.Context, page int, pageSize int) (list []v1.UserInfo, err error) {

	db := g.Model("users")
	//查询数据
	result, err := db.Page(page, pageSize).All()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	//判断是否有数据
	if result != nil {
		//循环遍历数据
		for _, v := range result {
			var isBanned bool
			if v["is_banned"].String() == "1" {
				isBanned = true
			} else {
				isBanned = false
			}
			//获取数据
			list = append(list, v1.UserInfo{
				UserId:     v["id"].Int64(),
				UserName:   v["username"].String(),
				Email:      v["email"].String(),
				CreateTime: v["create_time"].String(),
				UpdateTime: v["update_time"].String(),
				IsBanned:   isBanned,
			})
		}

	}
	return
}

// GetUserInfo	获取用户信息
func GetUserInfo(ctx context.Context, userId int64) (userInfo v1.UserInfo, err error) {
	db := g.Model("users")
	//查询数据
	result, err := db.Where("id", userId).One()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	//判断是否有数据
	if result != nil {
		var isBanned bool
		if result["is_banned"].String() == "1" {
			isBanned = true
		} else {
			isBanned = false
		}
		//获取数据
		userInfo = v1.UserInfo{
			UserId:     result["id"].Int64(),
			UserName:   result["username"].String(),
			Email:      result["email"].String(),
			CreateTime: result["create_time"].String(),
			UpdateTime: result["update_time"].String(),
			IsBanned:   isBanned,
		}
	}
	return
}
