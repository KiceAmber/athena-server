// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Article is the golang structure for table article.
type Article struct {
	Id          int         `json:"id"          ` // 文章ID
	Title       string      `json:"title"       ` // 文章标题
	Content     string      `json:"content"     ` // 文章内容
	Cover       string      `json:"cover"       ` // 文章图片封面
	AuthorId    int         `json:"author_id"   ` // 文章作者ID
	CategoryId  int         `json:"category_id" ` // 文章分类ID
	Description string      `json:"description" ` // 文章概述
	IsVisible   int         `json:"is_visible"  ` // 文章是否可见 1-可见 0-不可见
	CreatedAt   *gtime.Time `json:"created_at"  ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updated_at"  ` // 更新时间
	DeletedAt   *gtime.Time `json:"deleted_at"  ` // 删除时间
}
