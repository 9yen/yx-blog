package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	v1 "yx-blog/api/v1"
	"yx-blog/internal/consts"
	"yx-blog/internal/logic/article"
)

type cArticle struct {
}

var (
	Article = cArticle{}
)

// AddArticle 添加文章
func (c *cArticle) AddArticle(ctx context.Context, req *v1.AddArticleReq) (res *v1.AddArticleRes, err error) {
	//接收myKey
	myKey := req.MyKey
	//判断myKey是否正确
	if myKey != consts.MyKey {
		err = gerror.NewCode(gcode.New(consts.ErrMyKey, consts.ErrMyKeyMsg, nil))
		return res, err
	}

	//接收参数
	title := req.Title
	content := req.Content
	userId := req.UserId
	id, createTime := article.Add(ctx, title, content, userId)
	//返回数据
	res = &v1.AddArticleRes{
		ArticleId:  id,
		UserId:     userId,
		Title:      title,
		Content:    content,
		CreateTime: createTime,
	}
	return res, err
}

// DeleteArticle 删除文章
func (c *cArticle) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (res *v1.DeleteArticleRes, err error) {
	//接收myKey
	myKey := req.MyKey
	//判断myKey是否正确
	if myKey != consts.MyKey {
		err = gerror.NewCode(gcode.New(consts.ErrMyKey, consts.ErrMyKeyMsg, nil))
		return res, err
	}

	//接收参数
	articleId := req.ArticleId

	// check article exist
	exist := article.Check(ctx, articleId)

	if !exist {
		err = gerror.NewCode(gcode.New(consts.ErrArticleNotExist, consts.ErrArticleNotExistMsg, nil))
		return res, err
	}

	err = article.Delete(ctx, articleId)
	if err != nil {
		return nil, err
	}
	//返回数据
	res = &v1.DeleteArticleRes{
		ArticleId: articleId,
	}
	return res, err
}

// UpdateArticle 修改文章
func (c *cArticle) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error) {
	//接收myKey
	myKey := req.MyKey
	//判断myKey是否正确
	if myKey != consts.MyKey {
		err = gerror.NewCode(gcode.New(consts.ErrMyKey, consts.ErrMyKeyMsg, nil))
		return res, err
	}

	//接收参数
	articleId := req.ArticleId
	title := req.Title
	content := req.Content

	// check article exist
	exist := article.Check(ctx, articleId)

	if !exist {
		err = gerror.NewCode(gcode.New(consts.ErrArticleNotExist, consts.ErrArticleNotExistMsg, nil))
		return res, err
	}

	err, updateTime := article.Update(ctx, articleId, title, content)
	if err != nil {
		return nil, err
	}
	//返回数据
	res = &v1.UpdateArticleRes{
		ArticleId:  articleId,
		UpdateTime: updateTime,
	}
	return res, err
}
