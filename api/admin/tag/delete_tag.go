package tag

import "github.com/gogf/gf/v2/frame/g"

type DeleteTagReq struct {
	g.Meta `path:"/deleteTag" tags:"tag" method:"DELETE" sm:"删除已经存在的标签"`
	Id     int `json:"id"`
}

type DeleteTagRes struct {
	g.Meta `mime:"application/json"`
}
