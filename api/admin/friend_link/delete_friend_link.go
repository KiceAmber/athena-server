package friendlink

import "github.com/gogf/gf/v2/frame/g"

type DeleteFriendLinkReq struct {
	g.Meta `path:"/deleteFriendLink" method:"POST" tags:"friend_link" sm:"删除友链"`
	Id     int `json:"id"`
}

type DeleteFriendLinkRes struct{}
