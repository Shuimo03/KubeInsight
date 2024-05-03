package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"unique"` //保证用户名唯一
	Password  string    // 密码应当加密存储，这里暂时以明文形式存储
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	Roles     []Role    `gorm:"many2many:user_roles;"`
}
