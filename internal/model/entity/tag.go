// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tag is the golang structure for table tag.
type Tag struct {
	Id        int         `json:"id"         ` // 标签ID
	Name      string      `json:"name"       ` // 标签名称
	IsVisible int         `json:"is_visible" ` // 标签是否可见 1-可见 0-不可见
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" ` // 更新时间
	DeletedAt *gtime.Time `json:"deleted_at" ` // 删除时间
}
