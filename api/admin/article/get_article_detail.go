package article

import "github.com/gogf/gf/v2/frame/g"

type GetArticleDetailReq struct {
	g.Meta `path:"/getArticleDetail" method:"GET" tags:"article" sm:"查询单篇文章详情"`
	Id     int `json:"id"`
}

type GetArticleDetailRes struct {
	Article any `json:"article"`
}
