package admin

import "github.com/gogf/gf/v2/frame/g"

type GetTagListReq struct {
	g.Meta `path:"/getTagList" method:"get" tags:"tag" sm:"获取标签列表"`
}

type GetTagListRes struct {
	g.Meta  `mime:"application/json"`
	TagList any `json:"tagList"`
}

type CreateTagReq struct {
	g.Meta `path:"/createTag" tags:"tag" method:"POST" sm:"创建新的标签"`
	Name   string `json:"name"`
}

type CreateTagRes struct {
	g.Meta `mime:"application/json"`
}
