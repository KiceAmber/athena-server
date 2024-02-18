// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Category is the golang structure for table category.
type Category struct {
	Id        int         `json:"id"         ` // 分类ID
	Name      string      `json:"name"       ` // 分类名称
	AuthorId  int         `json:"author_id"  ` // 分类作者ID
	IsVisible int         `json:"is_visible" ` // 分类是否可见 1-可见 0-不可见
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" ` // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" ` // 删除时间
}
