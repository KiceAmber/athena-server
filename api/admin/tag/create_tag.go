package tag

import "github.com/gogf/gf/v2/frame/g"

type CreateTagReq struct {
	g.Meta  `path:"/createTag" tags:"tag" method:"POST" sm:"创建新的标签"`
	TagName string `json:"tag_name"`
}

type CreateTagRes struct {
}
