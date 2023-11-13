// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package dao

import (
	"time"
)

// 用户表
type User struct {
	// 主键id
	ID int64
	// 用户的名称
	Username string
	// 用户状态，1=正常，2=禁用
	Status int32
	// 邮箱
	Email string
	// 加密后的密码
	Password string
	// 备注
	Description string
	// 创建时间
	CreatedAt time.Time
	// 更新时间
	UpdatedAt time.Time
}
