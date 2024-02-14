package tag

import "github.com/gogf/gf/v2/frame/g"

type AddTagReq struct {
	g.Meta    `path:"/addTag" tags:"tag" method:"POST" sm:"添加新的标签"`
	Name      string `json:"name"`
	IsVisible int    `json:"isVisible"`
}

type AddTagRes struct {
	g.Meta `mime:"application/json"`
}
