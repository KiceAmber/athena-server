package category

import "github.com/gogf/gf/v2/frame/g"

type DeleteCategoryReq struct {
	g.Meta `path:"/deleteCategory" method:"POST" tags:"category" sm:"删除分类"`
	Id     int `json:"id"`
}

type DeleteCategoryRes struct{}
