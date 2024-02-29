package article

import "github.com/gogf/gf/v2/frame/g"

type GetArticleDetailReq struct {
	g.Meta `path:"/:id" method:"GET" tags:"article" sm:"获取文章详情"`
}

type GetArticleDetailRes struct {
	Content string `json:"content"`
}
