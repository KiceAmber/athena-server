package tag

import "github.com/gogf/gf/v2/frame/g"

type UpdateTagReq struct {
	g.Meta    `path:"/updateTag" tags:"tag" method:"POST" sm:"更新标签信息"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsVisible int    `json:"isVisible"`
}

type UpdateTagRes struct {
	g.Meta `mime:"application/json"`
}
