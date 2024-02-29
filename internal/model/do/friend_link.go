// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FriendLink is the golang structure of table friend_link for DAO operations like Where/Data.
type FriendLink struct {
	g.Meta    `orm:"table:friend_link, do:true"`
	Id        interface{} // 友链ID
	Url       interface{} // 链接路径
	Icon      interface{} // 友链图标
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
