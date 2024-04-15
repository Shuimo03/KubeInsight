package model

type Permission struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string // 权限名称
	Description string // 权限描述
	Roles       []Role `gorm:"many2many:role_permissions;"`
}
