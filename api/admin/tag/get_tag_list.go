package tag

import "github.com/gogf/gf/v2/frame/g"

type GetTagListReq struct {
	g.Meta `path:"/getTagList" method:"GET" tags:"tag" sm:"获取标签列表"`
}

type GetTagListRes struct {
	g.Meta  `mime:"application/json"`
	TagList any `json:"tagList"`
}
