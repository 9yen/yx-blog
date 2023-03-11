package article

import (
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Add is the logic method for adding article
func Add(ctx context.Context, title string, content string, userId int64) (id int64, createTime string) {
	db := g.Model("articles")

	//获取当前时间
	createTime = gtime.Now().String()

	//插入数据
	result, err := db.Insert(g.Map{
		"title":       title,
		"content":     content,
		"user_id":     userId,
		"create_time": createTime,
		"update_time": createTime,
	})

	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}

	//把result转换为int64
	id, err = result.LastInsertId()

	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}

	return id, createTime
}

// Check is checking article is exist
func Check(ctx context.Context, id int64) (exist bool) {
	db := g.Model("articles")
	//查询数据
	data, err := db.Where("id", id).One()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}

	if data != nil {
		exist = true
	}

	return
}

// Delete is the logic method for deleting article
func Delete(ctx context.Context, id int64) (err error) {
	db := g.Model("articles")
	//删除数据
	_, err = db.Where("id", id).Delete()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	return
}

// Update is the logic method for updating article
func Update(ctx context.Context, id int64, title string, content string) (err error, updateTime string) {
	db := g.Model("articles")
	updateTime = gtime.Now().String()

	//更新数据
	_, err = db.Data(g.Map{
		"title":       title,
		"content":     content,
		"update_time": updateTime,
	}).Where("id", id).Update()
	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}
	return
}

// Get is the logic method for getting article
func Get(ctx context.Context, id int64) (data g.Map, err error) {
	db := g.Model("articles")
	//查询数据
	res, err := db.Where("id", id).One()

	//判断是否有错误
	if err != nil {
		//打印错误
		g.Log().Error(ctx, err)
	}

	// res转换为map
	res.Map()

	//返回数据
	return data, err

}
