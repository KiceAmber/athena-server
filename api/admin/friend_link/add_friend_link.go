package friendlink

import "github.com/gogf/gf/v2/frame/g"

type AddFriendLinkReq struct {
	g.Meta `path:"/addFriendLink" method:"POST" tags:"friend_link" sm:"添加友链"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Url    string `json:"url"`
}

type AddFriendLinkRes struct{}
