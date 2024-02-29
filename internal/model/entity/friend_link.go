// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FriendLink is the golang structure for table friend_link.
type FriendLink struct {
	Id        int         `json:"id"         ` // 友链ID
	Url       string      `json:"url"        ` // 链接路径
	Icon      string      `json:"icon"       ` // 友链图标
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" ` // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" ` // 删除时间
}
