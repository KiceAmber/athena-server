package article

import "github.com/gogf/gf/v2/frame/g"

type GetArticleListReq struct {
	g.Meta `path:"/getArticleList" method:"GET" tags:"article" sm:"获取文章列表"`
	Page   int
}

type GetArticleListRes struct {
	ArticleList any `json:"articleList"`
}
