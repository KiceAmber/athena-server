// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure of table tag for DAO operations like Where/Data.
type Tag struct {
	g.Meta    `orm:"table:tag, do:true"`
	Id        interface{} // 标签ID
	Name      interface{} // 标签名称
	AuthorId  interface{} // 创建标签的用户ID
	IsVisible interface{} // 标签是否可见 1-可见 0-不可见
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
