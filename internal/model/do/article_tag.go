// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleTag is the golang structure of table article_tag for DAO operations like Where/Data.
type ArticleTag struct {
	g.Meta    `orm:"table:article_tag, do:true"`
	Id        interface{} // 文章标签关联ID
	ArticleId interface{} // 文章ID
	TagId     interface{} // 标签ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
