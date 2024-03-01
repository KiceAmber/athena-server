package category

import "github.com/gogf/gf/v2/frame/g"

type UpdateCategoryReq struct {
	g.Meta    `path:"/updateCategory" method:"POST" tags:"category" sm:"更新分类"`
	Id        int    `json:"id"`
	IsVisible int    `json:"isVisible"`
	Name      string `json:"name"`
}

type UpdateCategoryRes struct {
}
