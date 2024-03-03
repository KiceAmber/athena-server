package friendlink

import "github.com/gogf/gf/v2/frame/g"

type UpdateFriendLinkReq struct {
	g.Meta `path:"/updateFriendLink" method:"POST" tags:"friend_link" sm:"更新友链"`
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Url    string `json:"url"`
	Icon   string `json:"icon"`
}

type UpdateFriendLinkRes struct {
}
