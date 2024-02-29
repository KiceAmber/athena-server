// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure of table article for DAO operations like Where/Data.
type Article struct {
	g.Meta      `orm:"table:article, do:true"`
	Id          interface{} // 文章ID
	Title       interface{} // 文章标题
	Content     interface{} // 文章内容
	Cover       interface{} // 文章图片封面
	AuthorId    interface{} // 文章作者ID
	CategoryId  interface{} // 文章分类ID
	Description interface{} // 文章概述
	IsVisible   interface{} // 文章是否可见 1-可见 0-不可见
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
