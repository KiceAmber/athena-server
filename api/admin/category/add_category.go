package category

import "github.com/gogf/gf/v2/frame/g"

type AddCategoryReq struct {
	g.Meta    `path:"/addCategory" method:"POST" tags:"category" sm:"添加分类"`
	Name      string `json:"name"`
	IsVisible int    `json:"isVisible"`
}

type AddCategoryRes struct {
}
