// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        int         `json:"id"         ` // 用户ID
	Passport  string      `json:"passport"   ` // 用户账户
	Password  string      `json:"password"   ` // 用户密码
	Email     string      `json:"email"      ` // 用户邮箱
	Signature string      `json:"signature"  ` // 用户个人简介
	CreatedAt *gtime.Time `json:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updated_at" ` // 修改时间
	DeletedAt *gtime.Time `json:"deleted_at" ` // 删除时间
}
