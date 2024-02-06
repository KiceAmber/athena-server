package article

import "github.com/gogf/gf/v2/frame/g"

type AddArticleReq struct {
	g.Meta  `path:"/addArticle" method:"POST" tags:"article" sm:"添加文章"`
	Content string `json:"content"`
}

type AddArticleRes struct {
}
