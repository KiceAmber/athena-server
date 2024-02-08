// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ArticleTag is the golang structure for table article_tag.
type ArticleTag struct {
	Id        int         `json:"id"         ` // 文章标签关联ID
	ArticleId int         `json:"article_id" ` // 文章ID
	TagId     int         `json:"tag_id"     ` // 标签ID
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" ` // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" ` // 删除时间
}
