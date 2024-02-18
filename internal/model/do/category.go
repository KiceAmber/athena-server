// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure of table category for DAO operations like Where/Data.
type Category struct {
	g.Meta    `orm:"table:category, do:true"`
	Id        interface{} // 分类ID
	Name      interface{} // 分类名称
	AuthorId  interface{} // 分类作者ID
	IsVisible interface{} // 分类是否可见 1-可见 0-不可见
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
