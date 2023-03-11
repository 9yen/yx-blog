package v1

import "github.com/gogf/gf/v2/frame/g"

// AddArticleReq is the request struct for api AddArticle
type AddArticleReq struct {
	g.Meta  `path:"/article/add" method:"post" summary:"添加文章" tags:"文章"`
	UserId  int64  `p:"userId"  v:"required#用户ID不能为空"`
	Title   string `p:"title"  v:"required#文章标题不能为空"`
	Content string `p:"content"  v:"required#文章内容不能为空"`
	MyKey   string `p:"myKey"  v:"required#myKey不能为空"`
}

// AddArticleRes is the response struct for api AddArticle
type AddArticleRes struct {
	ArticleId  int64  `json:"articleId"`
	UserId     int64  `json:"userId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

// DeleteArticleReq is the request struct for api DeleteArticle
type DeleteArticleReq struct {
	g.Meta    `path:"/article/delete" method:"post" summary:"删除文章" tags:"文章"`
	ArticleId int64  `p:"articleId"  v:"required#文章ID不能为空"`
	MyKey     string `p:"myKey"  v:"required#myKey不能为空"`
}

type DeleteArticleRes struct {
	ArticleId int64 `json:"articleId"`
}
