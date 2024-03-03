package article

import "github.com/gogf/gf/v2/frame/g"

type GetArticleDetailReq struct {
	g.Meta `path:"/getArticleDetail" method:"GET" tags:"article" sm:"添加文章"`
}
