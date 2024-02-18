package category

import "github.com/gogf/gf/v2/frame/g"

type GetCategoryListReq struct {
	g.Meta `path:"/getCategoryList" method:"GET" tags:"category" sm:"获取分类列表"`
}

type GetCategoryListRes struct {
	CategoryList any `json:"categoryList"`
}
