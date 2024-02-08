package article

import "github.com/gogf/gf/v2/frame/g"

type DeleteArticleReq struct {
	g.Meta `path:"/deleteArticle" method:"DELETE" tags:"article" sm:"删除文章"`
	Id     int `json:"id"`
}

type DeleteArticleRes struct {
}
