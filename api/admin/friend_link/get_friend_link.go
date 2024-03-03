package friendlink

import "github.com/gogf/gf/v2/frame/g"

type GetFriendLinkListReq struct {
	g.Meta `path:"/getFriendLinkList" method:"GET" tags:"friend_link" sm:"获取友链列表"`
}

type GetFriendLinkListRes struct {
	FriendLinkList any `json:"friendLinkList"`
}
